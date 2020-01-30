package main

import (
	"errors"
	"fmt"

	"github.com/KnicKnic/go-powershell/pkg/powershell"
)

func executePowershell(script string, namedArgs map[string]interface{}) (obj []string, err error) {
	// runspace := powershell.CreateRunspace(logger.SimpleFuncPtr{func(line string) { fmt.Println(line) }}, nil)
	runspace := powershell.CreateRunspace(nil, nil)

	defer runspace.Close()

	results := runspace.ExecScript(script, true, namedArgs)
	defer results.Close()

	if !results.Success() {
		exception := results.Exception.ToString()
		fmt.Println("Error ", exception)
		return []string{}, errors.New(exception)
	}

	strResults := make([]string, len(results.Objects), len(results.Objects))
	for i, result := range results.Objects {
		strResults[i] = result.ToString()
	}

	return strResults, nil
}

func ensureVip(name, network, group, clusterName string) (string, error) {
	command := `
param(
	$resName,
	$groupName,
	$clusterNetwork,
	$cluster
)
$ErrorActionPreference = [System.Management.Automation.ActionPreference]::Stop

# $resName = "default%example-lb"
# $groupName = "g2"
# $clusterNetwork = "Cluster Network 2"
# $cluster = "localhost"
# $vipToClone = "Cluster IP Address"

# $cVip = get-clusterresource $vipToClone -cluster $cluster
# $cVipParam = $cVip | get-clusterparameter -cluster $cluster
write-information $resName
write-information $cluster
write-information $clusterNetwork 
write-information $groupName 
try{
	$r = Add-ClusterResource $resName -cluster $cluster -group $groupName -Type "IP Address" -ErrorAction SilentlyContinue
	}
	catch{
	
	}
$r = get-clusterresource $resName -cluster $cluster
$r.description = "Load Balancer"
#cluster will fix up the subnet :-)
$r |set-clusterparameter  -cluster $cluster  -Multiple @{"Network"=$clusterNetwork;"EnableDhcp"=1} -ErrorAction SilentlyContinue

$r = $r | start-clusterresource -wait 30 -cluster $cluster
$r = $r | stop-clusterresource -wait 30 -cluster $cluster


($r | get-clusterparameter -cluster $cluster -name "Address").Value

`

	addresses, err := executePowershell(command, map[string]interface{}{
		"resName":        name,
		"groupName":      group,
		"clusterNetwork": network,
		"cluster":        clusterName,
	})

	if err != nil {
		return "", err
	}
	// fmt.Println("got from powershell", addresses)
	return addresses[0], nil
}

func listVipNames(group, clusterName string) ([]string, error) {
	command := `
param(
	$groupName,
	$cluster
)
$ErrorActionPreference = [System.Management.Automation.ActionPreference]::Stop

Get-ClusterResource  -cluster $cluster |?{$_.ResourceType -eq "IP Address" -and $_.description -eq "Load Balancer"}|%{$_.name}

`

	names, err := executePowershell(command, map[string]interface{}{
		"groupName": group,
		"cluster":   clusterName,
	})

	if err != nil {
		return []string{}, err
	}
	return names, nil
}

func deleteVip(name, clusterName string) error {
	command := `
param(
	$name,
	$cluster
)
$ErrorActionPreference = [System.Management.Automation.ActionPreference]::Stop

Remove-ClusterResource  -cluster $cluster $name -force

`

	_, err := executePowershell(command, map[string]interface{}{
		"name":    name,
		"cluster": clusterName,
	})

	return err
}

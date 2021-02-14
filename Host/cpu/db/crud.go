package db

import (
	"cpu/status"
)

func HostCheckAll() ([]status.HostList, error) {
	db, err := DbCommect()
	if err != nil {
		return nil, err
	}
	var hostlist []status.HostList
	err = db.Select("id", "hostname", "ip_add", "port").Find(&hostlist).Error
	if err != nil {
		return nil, err
	}
	return hostlist, nil
}

func HostCheckOne(id int) (status.HostList, error) {
	db, err := DbCommect()
	if err != nil {
		return status.HostList{}, err
	}
	var hostlist status.HostList
	err = db.First(&hostlist, id).Error
	if err != nil {
		return status.HostList{}, err
	}
	return hostlist, nil
}

func CpuCheckOne(hostname string) (status.CPUStat, error) {
	db, err := DbCommect()
	if err != nil {
		return status.CPUStat{}, err
	}
	var cpustat status.CPUStat
	err = db.Where(&status.CPUStat{Hostname: hostname}).Find(&cpustat).Error
	if err != nil {
		return status.CPUStat{}, err
	}
	return cpustat, err
}

func VirCheckOne(hostname string) (status.VirtualMemoryStat, error) {
	db, err := DbCommect()
	if err != nil {
		return status.VirtualMemoryStat{}, err
	}
	var virstat status.VirtualMemoryStat
	err = db.Where(&status.VirtualMemoryStat{Hostname: hostname}).Find(&virstat).Error
	if err != nil {
		return status.VirtualMemoryStat{}, err
	}
	return virstat, err
}

func SwaCheckOne(hostname string) (status.SwapMemoryStat, error) {
	db, err := DbCommect()
	if err != nil {
		return status.SwapMemoryStat{}, err
	}
	var swastat status.SwapMemoryStat
	err = db.Where(&status.SwapMemoryStat{Hostname: hostname}).Find(&swastat).Error
	if err != nil {
		return status.SwapMemoryStat{}, err
	}
	return swastat, err
}

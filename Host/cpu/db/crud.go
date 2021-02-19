package db

import (
	"cpu/status"
	"fmt"
)

func HostCheckAll() ([]status.HostList, error) {
	db, err := DbCommect()
	if err != nil {
		return nil, err
	}
	var hostlist []status.HostList
	dbClose, err := db.DB()
	if err != nil {
		fmt.Println(err)
		return hostlist, err
	}
	defer dbClose.Close()
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
	dbClose, err := db.DB()
	if err != nil {
		fmt.Println(err)
		return hostlist, nil
	}
	defer dbClose.Close()
	err = db.First(&hostlist, id).Error
	if err != nil {
		return status.HostList{}, err
	}
	return hostlist, nil
}

func CpuCheckOne(hostname string) ([]status.CPUStat, error) {
	db, err := DbCommect()
	if err != nil {
		return nil, err
	}
	var cpustat []status.CPUStat
	dbClose, err := db.DB()
	if err != nil {
		fmt.Println(err)
		return cpustat, nil
	}
	defer dbClose.Close()
	err = db.Where("hostname = ?", hostname).Order("id desc").Limit(5).Find(&cpustat).Error
	if err != nil {
		return nil, err
	}
	return cpustat, nil
}

func VirCheckOne(hostname string) ([]status.VirtualMemoryStat, error) {
	db, err := DbCommect()
	if err != nil {
		return nil, err
	}
	var virstat []status.VirtualMemoryStat
	dbClose, err := db.DB()
	if err != nil {
		fmt.Println(err)
		return virstat, nil
	}
	defer dbClose.Close()
	err = db.Where("hostname = ?", hostname).Order("id desc").Limit(5).Find(&virstat).Error
	if err != nil {
		return nil, err
	}
	return virstat, nil
}

func SwaCheckOne(hostname string) ([]status.SwapMemoryStat, error) {
	db, err := DbCommect()
	if err != nil {
		return nil, err
	}
	var swastat []status.SwapMemoryStat
	dbClose, err := db.DB()
	if err != nil {
		fmt.Println(err)
		return swastat, nil
	}
	defer dbClose.Close()
	err = db.Where("hostname = ?", hostname).Order("id desc").Limit(5).Find(&swastat).Error
	if err != nil {
		return nil, err
	}
	return swastat, nil
}

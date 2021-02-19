package db

import (
	"cpu/status"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/* DBに接続する */
func DbCommect() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	USER := os.Getenv("USER")
	PASS := os.Getenv("DBPASS")
	PROTOCOL := os.Getenv("DBIP")
	TABLE := os.Getenv("DBTABLE")

	connect := USER + ":" + PASS + "@" + PROTOCOL + "/" + TABLE + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(connect), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}

/* DBのマイグレーション */
func DbInit() {
	db, err := DbCommect()
	if err != nil {
		fmt.Println(err)
		return
	}
	dbClose, err := db.DB()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dbClose.Close()
	db.AutoMigrate(&status.HostInfoStat{})
	db.AutoMigrate(&status.VirtualMemoryStat{})
	db.AutoMigrate(&status.SwapMemoryStat{})
	db.AutoMigrate(&status.CPUStat{})
	db.AutoMigrate(&status.HostList{})
}

/* DBにinsert */
func DbInsert(data status.AllInfo) {
	db, err := DbCommect()
	if err != nil {
		fmt.Println(err)
		return
	}
	dbClose, err := db.DB()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dbClose.Close()

	host := data.Host
	cpu := data.CPU
	vir := data.VirMem
	swap := data.SwapMem
	var hostname status.HostList

	db.Create(&status.HostInfoStat{
		Hostname:   host.Hostname,
		IpAdd:      host.IpAdd,
		Port:       host.Port,
		Uptime:     host.Uptime,
		Procs:      host.Procs,
		OS:         host.OS,
		Platform:   host.Platform,
		KernelArch: host.KernelArch,
	})

	db.Create(&status.CPUStat{
		Hostname: cpu.Hostname,
		Cpu0:     cpu.Cpu0,
		Cpu1:     cpu.Cpu1,
		Cpu2:     cpu.Cpu2,
		Cpu3:     cpu.Cpu3,
	})

	db.Create(&status.VirtualMemoryStat{
		Hostname:    vir.Hostname,
		Total:       vir.Total,
		Available:   vir.Available,
		Used:        vir.Used,
		UsedPercent: vir.UsedPercent,
		Free:        vir.Free,
		Active:      vir.Active,
		Buffers:     vir.Buffers,
		Cached:      vir.Cached,
	})

	db.Create(&status.SwapMemoryStat{
		Hostname:    swap.Hostname,
		Total:       swap.Total,
		Used:        swap.Used,
		Free:        swap.Free,
		UsedPercent: swap.UsedPercent,
	})

	/* hostnameの重複を確認→hostnameの唯一性 */
	if host.Hostname == "" {
		return
	}
	if err = db.Where("hostname = ?", host.Hostname).First(&hostname).Error; err == nil {
		//fmt.Println("同じhostnameが使用されています")
		return
		/* 重複されてなければ追加 */
	} else {
		db.Create(&status.HostList{
			Hostname: host.Hostname,
			IpAdd:    host.IpAdd,
			Port:     host.Port,
			Forward:  host.Forward,
		})
	}
	return
}

//https://www.dbonline.jp/mysql/type/index8.html これは？

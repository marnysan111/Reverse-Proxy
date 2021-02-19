package manage

import (
	"cpu/db"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	hostdata, err := db.HostCheckAll()
	if err != nil {
		fmt.Println("HostCHeck ERROR:", err)
		return
	}
	c.HTML(http.StatusOK, "index.html", gin.H{"hostdata": hostdata})
}

func Detail(c *gin.Context) {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	host, err := db.HostCheckOne(id)

	if err != nil {
		fmt.Println("HostCheckOne ERROR:", err)
		return
	}
	cpu, err := db.CpuCheckOne(host.Hostname)
	if err != nil {
		fmt.Println("CpuCheckOne ERROR:", err)
		return
	}
	vir, err := db.VirCheckOne(host.Hostname)
	if err != nil {
		fmt.Println("VirCheckOne ERROR:", err)
		return
	}
	swa, err := db.SwaCheckOne(host.Hostname)
	if err != nil {
		fmt.Println("SwaCheckOne ERROR:", err)
		return
	}

	c.HTML(http.StatusOK, "detail.html", gin.H{"host": host, "cpu": cpu, "vir": vir, "swa": swa})
}

package main

import "fmt"

func main() {

	fmt.Println("创建虚拟机")
	machine_name := keystring("取一个虚拟机的名字?", "myvirtualbox")

	ostypes := []string{"Ubuntu_64", "WindowsXP", "Other", "Other_64", "Windows95", "Windows98", "WindowsXP_64", "Windows2003", "Windows2003_64",
		"Windows2008", "Windows2008_64", "Windows7", "Windows7_64", "Windows8", "Windows8_64", "Windows10", "Windows10_64",
		"Debian", "Debian_64", "RedHat", "RedHat_64", "Ubuntu"}
	ostype_name, _ := chooseOne("操作系统的类型，请输入提示数字: (0-"+fmt.Sprint(len(ostypes)-1)+")", ostypes)

	fmt.Println("添加磁盘控制器")
	ctrltypes := []string{"SATA", "IDE"}
	ctrltype, _ := chooseOne("选择磁盘控制器类型: (0-"+fmt.Sprint(len(ctrltypes)-1)+")", ctrltypes)
	diskctrl_name := ""
	if ctrltype == "SATA" {
		diskctrl_name = keystring("取一个SATA控制器的名字?", "SATAController")
	} else {
		diskctrl_name = keystring("取一个IDE控制器的名字?", "IDE Controller")
	}

	fmt.Println("创建一个虚拟磁盘")
	diskfile_name := keystring("虚拟磁盘文件的名字?", machine_name)
	disk_size := keystring("虚拟磁盘大小(MB)?", "10000")
	disk_fix_opt := yesorno("动态增大磁盘文件大小?", "y")

	fmt.Println("虚拟硬盘放入磁盘控制器")
	iso_path := keystring("光盘iso的路径?", "")

	fmt.Println("网络配置")
	bridgeadp_name := ""
	if yesorno("是否改为桥接模式?", "y") == "y" {
		bridgeadp_name = keystring("what is bridgeadp_name?", "eth1")
	}

	fmt.Println("调整内存大小")
	mem_size := keystring("内存大小(MB)?", "512")
	fmt.Println("调整cpu数量和运行峰值")
	cpu_num := keystring("cpu数量?(1-4)", "1")
	cpu_cap := keystring("cpu运行峰值?(1-100)", "80")

	mstsc_port := ""
	if yesorno("是否打开远程?", "y") == "y" {
		mstsc_port = keystring("端口号?", "11010")
	}

	line1 := "vboxmanage createvm --name " + machine_name + " --ostype " + ostype_name + " --register"
	line2 := ""
	if ctrltype == "SATA" {
		line2 = "vboxmanage storagectl  " + machine_name + " --name \"" + diskctrl_name + "\" --add sata --hostiocache on --bootable on"
	} else {
		line2 = "vboxmanage storagectl  " + machine_name + " --name \"" + diskctrl_name + "\" --add ide --controller PIIX4 --hostiocache on --bootable on"
	}
	line3 := "vboxmanage createhd --filename " + diskfile_name + ".vdi --size " + disk_size
	if disk_fix_opt == "n" {
		line3 = line3 + " --variant fixed"
	}
	line4 := "vboxmanage storageattach  " + machine_name + " --storagectl \"" + diskctrl_name + "\" --port 0 --device 0 --type hdd --medium /home/wayne/VirtualBox\\ VMs/" + machine_name + "/" + diskfile_name + ".vdi"
	line5 := "vboxmanage storageattach  " + machine_name + " --storagectl \"" + diskctrl_name + "\" --port 1 --device 0 --type dvddrive --medium " + iso_path
	line6 := ""
	line7 := ""
	if bridgeadp_name != "" {
		line6 = "vboxmanage modifyvm  " + machine_name + " --nic1 bridged"
		line7 = "vboxmanage modifyvm  " + machine_name + " --bridgeadapter1 " + bridgeadp_name
	}
	line8 := "VBoxManage modifyvm " + machine_name + " --memory \"" + mem_size + "\""
	line9 := "VBoxManage modifyvm " + machine_name + " --cpus " + cpu_num + " --cpuexecutioncap " + cpu_cap
	line10 := ""
	if mstsc_port != "" {
		line10 = "vboxmanage modifyvm " + machine_name + " --vrde on --vrdeport " + mstsc_port
	}

	fmt.Println(line1)
	fmt.Println(line2)
	fmt.Println(line3)
	fmt.Println(line4)
	fmt.Println(line5)
	fmt.Println(line6)
	fmt.Println(line7)
	fmt.Println(line8)
	fmt.Println(line9)
	fmt.Println(line10)

	fmt.Println(machine_name, ostype_name, ctrltype, diskctrl_name, diskfile_name, disk_size, disk_fix_opt, iso_path, bridgeadp_name, mem_size, cpu_cap, cpu_num, mstsc_port)
}

func chooseOne(tishi string, xuanze []string) (string, int) {
	var ans int
	fmt.Println(tishi, "--默认: (0)")
	fmt.Println()
	for k, v := range xuanze {
		fmt.Println(fmt.Sprint(k) + "\t" + v)
	}
	fmt.Scanln(&ans)
	// if ans == nil {
	// 	return xuanze[0], 0
	// }
	if ans < 0 || ans > len(xuanze)-1 {
		return chooseOne(tishi+"\n(必须输入符合范围的数字)", xuanze)
	}
	return xuanze[ans], ans
}

func keystring(tishi string, moren string) string {
	var ans string
	fmt.Println(tishi, "--默认: ("+moren+")")
	fmt.Scanln(&ans)
	if ans == "" {
		return moren
	}
	return ans
}

func yesorno(tishi string, moren string) string {
	var ans string
	fmt.Print(tishi, "--默认: ("+moren+")   (y/n)")
	fmt.Scanln(&ans)
	if ans == "" {
		ans = moren
	}
	if !(ans == "y" || ans == "n") {
		return yesorno(tishi, moren)
	}
	return ans
}

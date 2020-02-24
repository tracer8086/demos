package main

import (
	"github.com/tracer8086/glfw/v3.3/glfw"
	"github.com/vulkan-go/demos/vulkaninfo"
	vk "github.com/vulkan-go/vulkan"
)

var appInfo = &vk.ApplicationInfo{
	SType:              vk.StructureTypeApplicationInfo,
	ApiVersion:         vk.MakeVersion(1, 0, 0),
	ApplicationVersion: vk.MakeVersion(1, 0, 0),
	PApplicationName:   "VulkanInfo\x00",
	PEngineName:        "vulkango.com\x00",
}

func main() {
	orPanic(glfw.Init())
	vk.SetGetInstanceProcAddr(glfw.GetVulkanGetInstanceProcAddress())
	orPanic(vk.Init())

	glfw.WindowHint(glfw.ClientAPI, glfw.NoAPI)
	window, err := glfw.CreateWindow(640, 480, "Vulkan Info", nil, nil)
	orPanic(err)
	defer window.Destroy()

	vkDevice, err := vulkaninfo.NewVulkanDevice(appInfo, window.GLFWWindow())
	orPanic(err)
	vulkaninfo.PrintInfo(vkDevice)
	vkDevice.Destroy()
}

func orPanic(err interface{}) {
	switch v := err.(type) {
	case error:
		if v != nil {
			panic(err)
		}
	case vk.Result:
		if err := vk.Error(v); err != nil {
			panic(err)
		}
	case bool:
		if !v {
			panic("condition failed: != true")
		}
	}
}

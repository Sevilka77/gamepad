package main

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

type DWORD uint32

const (
	XINPUT_DEVTYPE_GAMEPAD              = 0x01
	XINPUT_DEVSUBTYPE_GAMEPAD           = 0x01
	XINPUT_DEVSUBTYPE_WHEEL             = 0x02
	XINPUT_DEVSUBTYPE_ARCADE_STICK      = 0x03
	XINPUT_DEVSUBTYPE_FLIGHT_STICK      = 0x04
	XINPUT_DEVSUBTYPE_DANCE_PAD         = 0x05
	XINPUT_DEVSUBTYPE_GUITAR            = 0x06
	XINPUT_DEVSUBTYPE_DRUM_KIT          = 0.08
	XINPUT_CAPS_VOICE_SUPPORTED         = 0x0004
	XINPUT_GAMEPAD_DPAD_UP              = 0x0001
	XINPUT_GAMEPAD_DPAD_DOWN            = 0x0002
	XINPUT_GAMEPAD_DPAD_LEFT            = 0x0004
	XINPUT_GAMEPAD_DPAD_RIGHT           = 0x0008
	XINPUT_GAMEPAD_START                = 0x0010
	XINPUT_GAMEPAD_BACK                 = 0x0020
	XINPUT_GAMEPAD_LEFT_THUMB           = 0x0040
	XINPUT_GAMEPAD_RIGHT_THUMB          = 0x0080
	XINPUT_GAMEPAD_LEFT_SHOULDER        = 0x0100
	XINPUT_GAMEPAD_RIGHT_SHOULDER       = 0x0200
	XINPUT_GAMEPAD_A                    = 0x1000
	XINPUT_GAMEPAD_B                    = 0x2000
	XINPUT_GAMEPAD_X                    = 0x4000
	XINPUT_GAMEPAD_Y                    = 0x8000
	XINPUT_GAMEPAD_LEFT_THUMB_DEADZONE  = 7849
	XINPUT_GAMEPAD_RIGHT_THUMB_DEADZONE = 8689
	XINPUT_GAMEPAD_TRIGGER_THRESHOLD    = 30
	XINPUT_FLAG_GAMEPAD                 = 0x00000001
	BATTERY_DEVTYPE_GAMEPAD             = 0x00
	BATTERY_DEVTYPE_HEADSET             = 0x01
	BATTERY_TYPE_DISCONNECTED           = 0x00
	BATTERY_TYPE_WIRED                  = 0x01
	BATTERY_TYPE_ALKALINE               = 0x02
	BATTERY_TYPE_NIMH                   = 0x03
	BATTERY_TYPE_UNKNOWN                = 0xFF
	BATTERY_LEVEL_EMPTY                 = 0x00
	BATTERY_LEVEL_LOW                   = 0x01
	BATTERY_LEVEL_MEDIUM                = 0x02
	BATTERY_LEVEL_FULL                  = 0x03
	XINPUT_USER_0                       = DWORD(0)
	XINPUT_USER_1                       = DWORD(1)
	XINPUT_USER_2                       = DWORD(2)
	XINPUT_USER_3                       = DWORD(3)
	XUSER_MAX_COUNT                     = 4
	XUSER_INDEX_ANY                     = 0x000000FF
	XINPUT_GAMEPAD_TL_LIT               = DWORD(0)
	XINPUT_GAMEPAD_TR_LIT               = DWORD(1)
	XINPUT_GAMEPAD_BR_LIT               = DWORD(2)
	XINPUT_GAMEPAD_BL_LIT               = DWORD(3)
	VK_PAD_A                            = 0x5800
	VK_PAD_B                            = 0x5801
	VK_PAD_X                            = 0x5802
	VK_PAD_Y                            = 0x5803
	VK_PAD_RSHOULDER                    = 0x5804
	VK_PAD_LSHOULDER                    = 0x5805
	VK_PAD_LTRIGGER                     = 0x5806
	VK_PAD_RTRIGGER                     = 0x5807
	VK_PAD_DPAD_UP                      = 0x5810
	VK_PAD_DPAD_DOWN                    = 0x5811
	VK_PAD_DPAD_LEFT                    = 0x5812
	VK_PAD_DPAD_RIGHT                   = 0x5813
	VK_PAD_START                        = 0x5814
	VK_PAD_BACK                         = 0x5815
	VK_PAD_LTHUMB_PRESS                 = 0x5816
	VK_PAD_RTHUMB_PRESS                 = 0x5817
	VK_PAD_LTHUMB_UP                    = 0x5820
	VK_PAD_LTHUMB_DOWN                  = 0x5821
	VK_PAD_LTHUMB_RIGHT                 = 0x5822
	VK_PAD_LTHUMB_LEFT                  = 0x5823
	VK_PAD_LTHUMB_UPLEFT                = 0x5824
	VK_PAD_LTHUMB_UPRIGHT               = 0x5825
	VK_PAD_LTHUMB_DOWNRIGHT             = 0x5826
	VK_PAD_LTHUMB_DOWNLEFT              = 0x5827
	VK_PAD_RTHUMB_UP                    = 0x5830
	VK_PAD_RTHUMB_DOWN                  = 0x5831
	VK_PAD_RTHUMB_RIGHT                 = 0x5832
	VK_PAD_RTHUMB_LEFT                  = 0x5833
	VK_PAD_RTHUMB_UPLEFT                = 0x5834
	VK_PAD_RTHUMB_UPRIGHT               = 0x5835
	VK_PAD_RTHUMB_DOWNRIGHT             = 0x5836
	VK_PAD_RTHUMB_DOWNLEFT              = 0x5837
	XINPUT_KEYSTROKE_KEYDOWN            = 0x0001
	XINPUT_KEYSTROKE_KEYUP              = 0x0002
	XINPUT_KEYSTROKE_REPEAT             = 0x0004
	ERROR_SUCCESS                       = 0
	ERROR_DEVICE_NOT_CONNECTED          = 0x048F
	ERROR_EMPTY                         = 2
)

//State hold current state
type XINPUT_GAMEPAD struct {
	wButtons      uint16
	bLeftTrigger  uint8
	bRightTrigger uint8
	sThumbLX      uint16
	sThumbLY      uint16
	sThumbRX      uint16
	sThumbRY      uint16
}
type XINPUT_STATE struct {
	dwPacketNumber uint32
	Gamepad        XINPUT_GAMEPAD
}

type XINPUT_BATTERY_INFORMATION struct {
	BatteryType  uint8
	BatteryLevel uint8
}

type XINPUT_KEYSTROKE struct {
	VirtualKey uint16
	Unicode    uint16
	Flags      uint16
	UserIndex  uint8
	HidCode    uint8
}

type XINPUT_VIBRATION struct {
	wLeftMotorSpeed  uint16
	wRightMotorSpeed uint16
}

type XINPUT_CAPABILITIES struct {
	Type      uint8
	SubType   uint8
	Flags     uint16
	Gamepad   XINPUT_GAMEPAD
	Vibration XINPUT_VIBRATION
}
type Gamepad struct {
	id      int
	state   XINPUT_STATE
	battery XINPUT_BATTERY_INFORMATION
}

var (
	xinputdll                = windows.MustLoadDLL("Xinput1_4.dll")
	joyGetState              = xinputdll.MustFindProc("XInputGetState")
	joySetState              = xinputdll.MustFindProc("XInputSetState")
	joyGetCapabilities       = xinputdll.MustFindProc("XInputGetCapabilities")
	joyGetBatteryInformation = xinputdll.MustFindProc("XInputGetBatteryInformation")
)

func (js *Gamepad) Open(id int) error {
	js.id = id
	err := js.getJoyState()
	if err == nil {
		return nil
	}
	return err

}

func (js *Gamepad) getBatteryLevel() string {
	switch js.battery.BatteryLevel {
	case BATTERY_LEVEL_FULL:
		return "BATTERY_LEVEL_FULL"
	case BATTERY_LEVEL_MEDIUM:
		return "BATTERY_LEVEL_MEDIUM"
	case BATTERY_LEVEL_LOW:
		return "BATTERY_LEVEL_LOW"
	default:
		return "BATTERY_LEVEL_EMPTY"
	}
}
func (js *Gamepad) getBatteryType() string {
	switch js.battery.BatteryType {
	case BATTERY_TYPE_ALKALINE:
		return "BATTERY_TYPE_ALKALINE"
	case BATTERY_TYPE_NIMH:
		return "BATTERY_TYPE_NIMH"
	case BATTERY_TYPE_UNKNOWN:
		return "BATTERY_TYPE_UNKNOWN"
	case BATTERY_TYPE_WIRED:
		return "BATTERY_TYPE_WIRED"
	default:
		return "BATTERY_TYPE_DISCONNECTED"
	}
}
func (js *Gamepad) getJoyState() error {
	var state XINPUT_STATE
	ret, _, _ := joyGetState.Call(uintptr(js.id), uintptr(unsafe.Pointer(&state)))
	if ret != 0 {
		return fmt.Errorf("Failed to read Joystick %d", js.id)
	} else {
		js.state = state
	}
	return nil
}
func (js *Gamepad) getBatteryInformation() error {
	var btstate XINPUT_BATTERY_INFORMATION
	ret, _, _ := joyGetBatteryInformation.Call(uintptr(js.id), uintptr(XINPUT_DEVTYPE_GAMEPAD), uintptr(unsafe.Pointer(&btstate)))
	if ret != 0 {
		return fmt.Errorf("Failed to read battary Joystick %d", js.id)
	} else {
		js.battery = btstate
	}
	return nil
}

func main() {
	fmt.Println("Привет попоробуем считать чтолибо")
	gamepads := make([]Gamepad, 4)
	for i := range gamepads {
		gamepads[i].id = i
		gamepads[i].getJoyState()
		gamepads[i].getBatteryInformation()
		fmt.Println(gamepads[i])
	}
	fmt.Println(gamepads[0].getBatteryType())
}

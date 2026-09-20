/******************************************************************************
 * Radio Code Calculator API - error codes
 *
 * Version      : v1.1.6
 * Language     : Go
 * Author       : Bartosz Wójcik (support@pelock.com)
 * Homepage     : https://www.pelock.com
 *
 *****************************************************************************/

package radiocodecalculator

const (
	ErrorConnection             = -1
	ErrorSuccess                = 0
	ErrorInvalidInput           = 1
	ErrorInvalidCommand         = 2
	ErrorInvalidRadioModel      = 3
	ErrorInvalidSerialLength    = 4
	ErrorInvalidSerialPattern   = 5
	ErrorInvalidSerialNotSupported = 6
	ErrorInvalidExtraLength     = 7
	ErrorInvalidExtraPattern    = 8
	ErrorInvalidLicense         = 100
)

package testassets

import (
	"time"

	"github.com/tscolari/plag/parser"
)

const Log string = `
{"timestamp":"1485092992.139814854","source":"test-app","message":"test-app.function-4.start","log_level":1,"data":{"session":"3"}}
{"timestamp":"1485092992.140044212","source":"test-app","message":"test-app.function-4.start-to-sleep","log_level":1,"data":{"session":"3"}}
{"timestamp":"1485092992.140080690","source":"test-app","message":"test-app.function-4.finished-to-sleep","log_level":1,"data":{"session":"3"}}
{"timestamp":"1485092992.139838219","source":"test-app","message":"test-app.function-4.start","log_level":1,"data":{"session":"4"}}
{"timestamp":"1485092992.140119314","source":"test-app","message":"test-app.function-4.start-to-sleep","log_level":1,"data":{"session":"4"}}
{"timestamp":"1485092992.139848709","source":"test-app","message":"test-app.function-0.start","log_level":1,"data":{"session":"6"}}
{"timestamp":"1485092992.140167475","source":"test-app","message":"test-app.function-0.start-to-sleep","log_level":1,"data":{"session":"6"}}
{"timestamp":"1485092992.139834881","source":"test-app","message":"test-app.function-0.start","log_level":1,"data":{"session":"1"}}
{"timestamp":"1485092992.140216351","source":"test-app","message":"test-app.function-0.start-to-sleep","log_level":1,"data":{"session":"1"}}
{"timestamp":"1485092992.139839411","source":"test-app","message":"test-app.function-2.start","log_level":1,"data":{"session":"5"}}
{"timestamp":"1485092992.140252829","source":"test-app","message":"test-app.function-2.start-to-sleep","log_level":1,"data":{"session":"5"}}
{"timestamp":"1485092992.139858484","source":"test-app","message":"test-app.function-1.start","log_level":1,"data":{"session":"7"}}
{"timestamp":"1485092992.140296698","source":"test-app","message":"test-app.function-1.start-to-sleep","log_level":1,"data":{"session":"7"}}
{"timestamp":"1485092992.139864922","source":"test-app","message":"test-app.function-2.start","log_level":1,"data":{"session":"9"}}
{"timestamp":"1485092992.140336752","source":"test-app","message":"test-app.function-2.start-to-sleep","log_level":1,"data":{"session":"9"}}
{"timestamp":"1485092992.139867306","source":"test-app","message":"test-app.function-3.start","log_level":1,"data":{"session":"10"}}
{"timestamp":"1485092992.140373230","source":"test-app","message":"test-app.function-3.start-to-sleep","log_level":1,"data":{"session":"10"}}
{"timestamp":"1485092992.139861822","source":"test-app","message":"test-app.function-3.start","log_level":1,"data":{"session":"8"}}
{"timestamp":"1485092992.140414715","source":"test-app","message":"test-app.function-3.start-to-sleep","log_level":1,"data":{"session":"8"}}
{"timestamp":"1485092992.139875650","source":"test-app","message":"test-app.function-2.start","log_level":1,"data":{"session":"11"}}
{"timestamp":"1485092992.140447855","source":"test-app","message":"test-app.function-2.start-to-sleep","log_level":1,"data":{"session":"11"}}
{"timestamp":"1485092992.139880896","source":"test-app","message":"test-app.function-1.start","log_level":1,"data":{"session":"13"}}
{"timestamp":"1485092992.140481710","source":"test-app","message":"test-app.function-1.start-to-sleep","log_level":1,"data":{"session":"13"}}
{"timestamp":"1485092992.139877558","source":"test-app","message":"test-app.function-4.start","log_level":1,"data":{"session":"12"}}
{"timestamp":"1485092992.140516520","source":"test-app","message":"test-app.function-4.start-to-sleep","log_level":1,"data":{"session":"12"}}
{"timestamp":"1485092992.139891624","source":"test-app","message":"test-app.function-0.start","log_level":1,"data":{"session":"15"}}
{"timestamp":"1485092992.140550375","source":"test-app","message":"test-app.function-0.start-to-sleep","log_level":1,"data":{"session":"15"}}
{"timestamp":"1485092992.139884710","source":"test-app","message":"test-app.function-3.start","log_level":1,"data":{"session":"14"}}
{"timestamp":"1485092992.140589714","source":"test-app","message":"test-app.function-3.start-to-sleep","log_level":1,"data":{"session":"14"}}
{"timestamp":"1485092992.139912128","source":"test-app","message":"test-app.function-4.start","log_level":1,"data":{"session":"18"}}
{"timestamp":"1485092992.140639782","source":"test-app","message":"test-app.function-4.start-to-sleep","log_level":1,"data":{"session":"18"}}
{"timestamp":"1485092992.139907837","source":"test-app","message":"test-app.function-2.start","log_level":1,"data":{"session":"16"}}
{"timestamp":"1485092992.140672445","source":"test-app","message":"test-app.function-2.start-to-sleep","log_level":1,"data":{"session":"16"}}
{"timestamp":"1485092992.139920712","source":"test-app","message":"test-app.function-0.start","log_level":1,"data":{"session":"19"}}
{"timestamp":"1485092992.140709162","source":"test-app","message":"test-app.function-0.start-to-sleep","log_level":1,"data":{"session":"19"}}
{"timestamp":"1485092992.139925718","source":"test-app","message":"test-app.function-1.start","log_level":1,"data":{"session":"20"}}
{"timestamp":"1485092992.140740156","source":"test-app","message":"test-app.function-1.start-to-sleep","log_level":1,"data":{"session":"20"}}
{"timestamp":"1485092992.139908552","source":"test-app","message":"test-app.function-3.start","log_level":1,"data":{"session":"17"}}
{"timestamp":"1485092992.140770912","source":"test-app","message":"test-app.function-3.start-to-sleep","log_level":1,"data":{"session":"17"}}
{"timestamp":"1485092992.139804840","source":"test-app","message":"test-app.function-1.start","log_level":1,"data":{"session":"2"}}
{"timestamp":"1485092992.140810966","source":"test-app","message":"test-app.function-1.start-to-sleep","log_level":1,"data":{"session":"2"}}
{"timestamp":"1485092992.140097618","source":"test-app","message":"test-app.function-4.end","log_level":1,"data":{"session":"3"}}
{"timestamp":"1485092992.140314579","source":"test-app","message":"test-app.function-1.finished-to-sleep","log_level":1,"data":{"session":"7"}}
{"timestamp":"1485092992.140860558","source":"test-app","message":"test-app.function-1.end","log_level":1,"data":{"session":"7"}}
{"timestamp":"1485092992.140785694","source":"test-app","message":"test-app.function-3.finished-to-sleep","log_level":1,"data":{"session":"17"}}
{"timestamp":"1485092992.140896797","source":"test-app","message":"test-app.function-3.end","log_level":1,"data":{"session":"17"}}
{"timestamp":"1485092993.141621590","source":"test-app","message":"test-app.function-1.finished-to-sleep","log_level":1,"data":{"session":"2"}}
{"timestamp":"1485092993.141746998","source":"test-app","message":"test-app.function-1.end","log_level":1,"data":{"session":"2"}}
{"timestamp":"1485092993.141653538","source":"test-app","message":"test-app.function-0.finished-to-sleep","log_level":1,"data":{"session":"1"}}
{"timestamp":"1485092993.141844511","source":"test-app","message":"test-app.function-0.end","log_level":1,"data":{"session":"1"}}
{"timestamp":"1485092993.141662359","source":"test-app","message":"test-app.function-1.finished-to-sleep","log_level":1,"data":{"session":"20"}}
{"timestamp":"1485092993.141916990","source":"test-app","message":"test-app.function-1.end","log_level":1,"data":{"session":"20"}}
{"timestamp":"1485092993.141702890","source":"test-app","message":"test-app.function-4.finished-to-sleep","log_level":1,"data":{"session":"4"}}
{"timestamp":"1485092993.142002106","source":"test-app","message":"test-app.function-4.end","log_level":1,"data":{"session":"4"}}
{"timestamp":"1485092993.141621590","source":"test-app","message":"test-app.function-0.finished-to-sleep","log_level":1,"data":{"session":"6"}}
{"timestamp":"1485092993.142052412","source":"test-app","message":"test-app.function-0.end","log_level":1,"data":{"session":"6"}}
{"timestamp":"1485092995.142178297","source":"test-app","message":"test-app.function-2.finished-to-sleep","log_level":1,"data":{"session":"16"}}
{"timestamp":"1485092995.142283201","source":"test-app","message":"test-app.function-2.end","log_level":1,"data":{"session":"16"}}
{"timestamp":"1485092996.141746759","source":"test-app","message":"test-app.function-1.finished-to-sleep","log_level":1,"data":{"session":"13"}}
{"timestamp":"1485092996.141840935","source":"test-app","message":"test-app.function-1.end","log_level":1,"data":{"session":"13"}}
{"timestamp":"1485092996.141746759","source":"test-app","message":"test-app.function-0.finished-to-sleep","log_level":1,"data":{"session":"15"}}
{"timestamp":"1485092996.141939640","source":"test-app","message":"test-app.function-0.end","log_level":1,"data":{"session":"15"}}
{"timestamp":"1485092997.142184019","source":"test-app","message":"test-app.function-4.finished-to-sleep","log_level":1,"data":{"session":"18"}}
{"timestamp":"1485092997.142268896","source":"test-app","message":"test-app.function-4.end","log_level":1,"data":{"session":"18"}}
{"timestamp":"1485092998.141907692","source":"test-app","message":"test-app.function-3.finished-to-sleep","log_level":1,"data":{"session":"8"}}
{"timestamp":"1485092998.142008305","source":"test-app","message":"test-app.function-3.end","log_level":1,"data":{"session":"8"}}
{"timestamp":"1485092998.141907692","source":"test-app","message":"test-app.function-3.finished-to-sleep","log_level":1,"data":{"session":"14"}}
{"timestamp":"1485092998.142096758","source":"test-app","message":"test-app.function-3.end","log_level":1,"data":{"session":"14"}}
{"timestamp":"1485092999.141546249","source":"test-app","message":"test-app.function-4.finished-to-sleep","log_level":1,"data":{"session":"12"}}
{"timestamp":"1485092999.141662836","source":"test-app","message":"test-app.function-4.end","log_level":1,"data":{"session":"12"}}
{"timestamp":"1485092999.141546249","source":"test-app","message":"test-app.function-2.finished-to-sleep","log_level":1,"data":{"session":"5"}}
{"timestamp":"1485092999.141761065","source":"test-app","message":"test-app.function-2.end","log_level":1,"data":{"session":"5"}}
{"timestamp":"1485093000.140420198","source":"test-app","message":"test-app.function-3.finished-to-sleep","log_level":1,"data":{"session":"10"}}
{"timestamp":"1485093000.140520811","source":"test-app","message":"test-app.function-3.end","log_level":1,"data":{"session":"10"}}
{"timestamp":"1485093000.140456438","source":"test-app","message":"test-app.function-2.finished-to-sleep","log_level":1,"data":{"session":"9"}}
{"timestamp":"1485093000.140604734","source":"test-app","message":"test-app.function-2.end","log_level":1,"data":{"session":"9"}}
{"timestamp":"1485093000.141000032","source":"test-app","message":"test-app.function-0.finished-to-sleep","log_level":1,"data":{"session":"19"}}
{"timestamp":"1485093000.141100645","source":"test-app","message":"test-app.function-0.end","log_level":1,"data":{"session":"19"}}
{"timestamp":"1485093001.140837193","source":"test-app","message":"test-app.function-2.finished-to-sleep","log_level":1,"data":{"session":"11"}}
{"timestamp":"1485093001.140943766","source":"test-app","message":"test-app.function-2.end","log_level":1,"data":{"session":"11"}}
`

var DataPointsPerMessage map[string][]parser.Data = map[string][]parser.Data{
	"test-app.function-0": []parser.Data{
		parser.Data{Timestamp: time.Unix(0, 1485092992139848704), Value: 1002203648},
		parser.Data{Timestamp: time.Unix(0, 1485092992139834880), Value: 1002009600},
		parser.Data{Timestamp: time.Unix(0, 1485092992139891712), Value: 4002048000},
		parser.Data{Timestamp: time.Unix(0, 1485092992139920640), Value: 8001179904},
	},
	"test-app.function-1": []parser.Data{
		parser.Data{Timestamp: time.Unix(0, 1485092992139858432), Value: 1002240},
		parser.Data{Timestamp: time.Unix(0, 1485092992139880960), Value: 4001959936},
		parser.Data{Timestamp: time.Unix(0, 1485092992139925760), Value: 1001991168},
		parser.Data{Timestamp: time.Unix(0, 1485092992139804928), Value: 1001942016},
	},
	"test-app.function-2": []parser.Data{
		parser.Data{Timestamp: time.Unix(0, 1485092992139839488), Value: 7001921536},
		parser.Data{Timestamp: time.Unix(0, 1485092992139864832), Value: 8000739840},
		parser.Data{Timestamp: time.Unix(0, 1485092992139875584), Value: 9001068288},
		parser.Data{Timestamp: time.Unix(0, 1485092992139907840), Value: 3002375424},
	},
	"test-app.function-3": []parser.Data{
		parser.Data{Timestamp: time.Unix(0, 1485092992139867392), Value: 8000653312},
		parser.Data{Timestamp: time.Unix(0, 1485092992139861760), Value: 6002146560},
		parser.Data{Timestamp: time.Unix(0, 1485092992139884800), Value: 6002211840},
		parser.Data{Timestamp: time.Unix(0, 1485092992139908608), Value: 988160},
	},
	"test-app.function-4": []parser.Data{
		parser.Data{Timestamp: time.Unix(0, 1485092992139814912), Value: 282624},
		parser.Data{Timestamp: time.Unix(0, 1485092992139838208), Value: 1002163968},
		parser.Data{Timestamp: time.Unix(0, 1485092992139877632), Value: 7001785088},
		parser.Data{Timestamp: time.Unix(0, 1485092992139912192), Value: 5002356736},
	},
}

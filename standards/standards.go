package standards

/******************** Drive Standards ***********/

const STOP_SIGN_TIME = 0.0499
const YIELD_SIGN_TIME = 0.0159
const TRAFFIC_LIGHT_TIME = 0.2051
const CROSSWALK_TIME = 0.0092
const RRCROSSING_TIME = 0.0165
const ACCESS_GATE_TIME = 0.1805
const CREEP_TIME = 0.0208

// Drive Speed Matrix

//min distance, max distance, speed (ft/min)
type DSMPoint struct {
	Min   float64
	Max   float64
	Speed float64
}

// from study results pg 300: Drive Speed Matrix
var DSM = []DSMPoint{
	{0, 15, 0.00249},
	{15, 20, 0.00176},
	{20, 30, 0.00177},
	{30, 45, 0.00168},
	{45, 65, 0.00153},
	{65, 90, 0.00141},
	{90, 120, 0.00128},
	{120, 155, 0.00116},
	{155, 195, 0.00106},
	{195, 240, 0.00096},
	{240, 290, 0.00088},
	{290, 345, 0.00082},
	{345, 405, 0.00076},
	{405, 470, 0.00071},
	{470, 540, 0.00067},
	{540, 615, 0.00063},
	{615, 695, 0.0006},
	{695, 780, 0.00057},
	{780, 870, 0.00054},
	{870, 965, 0.00052},
	{965, 1065, 0.0005},
	{1065, 1170, 0.00048},
	{1170, 1280, 0.00046},
	{1280, 1395, 0.00045},
	{1395, 1515, 0.00044},
	{1515, 1640, 0.00043},
	{1640, 1770, 0.00041},
	{1770, 1905, 0.0004},
	{1905, 2045, 0.00039},
	{2045, 2190, 0.00038},
	{2190, 2340, 0.00038},
	{2340, 2495, 0.00037},
	{2495, 2655, 0.00036},
	{2655, 2820, 0.00036},
	{2820, 2990, 0.00035},
	{2990, 3165, 0.00035},
	{3165, 3345, 0.00034},
	{3345, 3530, 0.00034},
	{3530, 3720, 0.00033},
	{3720, 3915, 0.00033},
	{3915, 4115, 0.00032},
	{4115, 4320, 0.00032},
	{4320, 4530, 0.00032},
	{4530, 4745, 0.00032},
	{4745, 4965, 0.00032},
	{4965, 5190, 0.00031},
	{5191, 5280, 0.00029},
}

/******************** END Drive Standards ***********/

/******************* DoorTrip Standards ***************/

const PRELIM_DOOR_TRIP_TIME = 0.852
const GATHER_DOOR_FIXED_TIME = 0.205 //per trip
const GATHER_DOOR_VAR_TIME = 0.0812  //per piece
const DELIVER_MISC_TO_DOOR_TIME = 0.0854

/****************** Averaged Time Standards **************/

//these values are averaged out over all routes
const PARCEL_ACCEPTED_TIME = 0.1034
const PARCEL_ACCEPTED_PROB = 0.0002

const REG_CERT_ACCEPTED_TIME = 0.0944
const REG_CERT_ACCEPTED_PROB = 0.000094

const MONEY_ORDER_ACCEPTED_TIME = 2.1799
const MONEY_ORDER_ACCEPTED_PROB = 0.000009

/****** Accountables, COD, PostageDue Standards ***********/

// Office time
const PROCESS_ACCOUNTABLE_MAIL = 1.4745
const FIXED_ACCOUNTABLE_CAGE = 0.7418 //per day
const PROCESS_POSTAGE_DUE = 0.0345

// Street time

const DELIVER_ACCOUNTABLE_SIG_TIME = 0.9615
const DELIVER_COD_TIME = 1.6131
const DELIVER_CUSTOMS_DUE_TIME = 1.6618

const COLLECT_POSTAGE_DUE = 0.9127

/********************* Box Time Standards *****************/

const VERIFY_LETTER = .0116
const VERIFY_FLAT = .0142

type BoxType int

const (
	CURB BoxType = iota
	SIDEWALK
	OTHER
)

// rows are box type (CURB, SDWK, OTHER) col are bundling (1,2,or 3)
var RegBox = [][]float64{
	{.1483, .2027, .2572},
	{.1356, .1900, .2445},
	{.1346, .1890, .2434},
}

/************************ Boxholder Time Standards **********/

const BOXHOLDER_FLATS_TIME = 0.0646
const BOXHOLDER_LETTERS_TIME = 0.0303

/******************* Centralized Box Time Standards ********/

const COLLECTION_TIME = 0.0558

// RRECS standards
const UNIT_TIME_CBU = 0.3311
const UNIT_TIME_CENT = 0.3803
const UNIT_TIME_NPU = 0.0983

//DET is also called a VPO
const UNIT_TIME_DET = 0.0983
const PRESORT_TIME_DET = 0.1256

const BOX_TIME_DET = 0.2084

var CLUSTER_BUNDLE_BOX_TIME = []float64{
	0.0528, 0.1077, 0.1625,
}

/******************* Driveway Time ***********************/

// these are my simplifications
const LESS_THAN_A_POOL_TIME = 0.10
const POOL_TIME = 0.25
const FOOTBALL_FIELD_TIME = 0.51
const TWO_FOOTBALL_FIELD_TIME = 0.75
const QUARTER_MILE_TIME = 1.25
const HALF_MILE_TIME = 1.85

/******************* Dismount Time ***********************/

// RRECS Standards
const OTHER_DISMOUNT_PREP_TIME = 0.8931
const OTHER_ADD_TRIP_TIME = 0.2076

const SDWK_DISMOUNT_PREP_TIME = 0.6709
const SDWK_ADD_TRIP_TIME = 0.1544

const CBU_DISMOUNT_PREP_TIME = 0.8816
const CBU_ADD_TRIP_TIME = 0.1544

const CENT_DISMOUNT_PREP_TIME = 1.9112
const CENT_ADD_TRIP_TIME = 0.4409

const NPU_DISMOUNT_PREP_TIME = 1.9112
const NPU_ADD_TRIP_TIME = 0.4409

const VPO_DISMOUNT_PREP_TIME = 2.076
const VPO_ADD_TRIP_TIME = 0.4958

/********************** Bundling ****************************/
type Bundling int

const (
	ONE_BUNDLE Bundling = iota
	TWO_BUNDLE
	THREE_BUNDLE
)

/********************* Flats Time Standards ***********************/
const RAW_FLATS_TIME = 0.0862
const CR_FLATS_TIME = 0.0741
const WSS_FLATS_TIME = 0.0741
const CASE_DPS_FLATS_TIME = 0.0708

/******************** Markup time standards **************************/

const MARKUP_TIME = 0.4233
const MARKUP_BUNDLE_RATE = .00147

/********************* Letter time standards **************************/

// paid to case if less than threshold
const DPS_THRESHOLD = 400

const RAW_LETTERS_NEW = 0.0647
const WSS_LETTERS_NEW = 0.0364
const CASE_DPS_LETTERS = 0.0294

/********************** Misc Time Standards ******************/

const FORM3982_PARS_LABEL = 0.405
const SERVICE_BLUE_BOX = 1.7928

// Same scan time for SPM too
const ONE_STEP_RRECS_SCAN = 0.0795
const TWO_STEP_RRECS_SCAN = 0.0928

const RURAL_REACH_TIME = 5.2458
const SELL_STAMP_STOCK_TIME = 0.7140

const SCANNER_SETUP_TIME = 0.1882

const TRIP_REPORT_TIME = 0.491
const MOVE_TRAYS_STORAGE_TIME = 0.2643

/***************** Withdraw Time Standards ***********/

const GATHER_RANDOM_MAIL_TIME = 1.0517

const ACCESS_DPS_LETTER_TIME = 0.7689
const GATHER_DPS_LETTER_TRAY_TIME = 0.25

const ACCESS_DPS_FLAT_TIME = 0.1653
const GATHER_DPS_FLAT_TRAY_TIME = 0.205

/******************** Locked Pouch Time Standards ************/

const LOW_VOLUME_POUCH_STOP_TIME = 1.4787
const HIGH_VOLUME_POUCH_STOP_TIME = 5.7659

/*************************** Parcels Time Standards *******************/

// cart times
const PARCEL_PER_CART = 67
const GATHER_1ST_CART_TIME = 0.1774
const GATHER_ADD_CART_TIME = 0.4944

// organize times
const ORGANIZE_LARGE_TIME = 0.5157
const ORGANIZE_SMALL_MED_TIME = 0.2367

// delivery times
const DELIVER_MAILBOX_TIME = 0.1671
const DELIVER_LOCKER_TIME = 0.3589
const DELIVER_DOOR_TIME = 0.1954

/*********************** Parcel Size Standards ***********************/

// Comprehensive guide page 94
// medium vs small is based off of statistics
const SMALL_PARCEL_PERCENT = 0.2174
const MEDIUM_PARCEL_PERCENT = 0.7826

/**************************** Parcel Pickup Time Standards ************/

const PICKUP_PROCESS_FORM_TIME = 0.114
const PICKUP_EVENT_TIME = 1.1817
const FULL_SCAN_TIME = 0.1011
const PICKUP_ITEM_TIME = 0.1592
const PARTIAL_SCAN_TIME = 0.0223

/********************** Pieces per Tray ***************************/
const LETTERS_PER_TRAY = 425
const FLATS_PER_TRAY = 115
const SMALL_PARCELS_PER_TRAY = 60
const MEDIUM_PARCELS_PER_TRAY = 10

/********************** Reload Time Standards ****************************/

const FRONT_LOADED_MAIL_TRAYS = 3
const FRONT_LOADED_MED_PARCEL_TRAY = 2
const REACHABLE_LARGE_PARCEL = 6

const RELOAD_LARGE_PARCEL_TIME = 0.1895

const FIRST_TRAY_REAR_DOOR = 1.7544
const ADD_TRAY_REAR_DOOR = 0.2156

const FIRST_PACKAGE_REAR_DOOR = 1.2172
const ADD_PACKAGE_REAR_DOOR = 0.2156

var MAIL_TRAY_TABLE = [][]float64{

	{0.0000, 0.0000, 0.0000},
	{0.6957, 0.6957, 0.6957},
	{0.8689, 0.8689, 1.3914},
	{2.2820, 2.2820, 2.8045},
	{2.4023, 3.6951, 4.2176},
	{4.1567, 5.4495, 5.9720},
	{4.3723, 5.6651, 7.7264},
	{4.5879, 7.4195, 9.4808},
	{6.3423, 7.6351, 11.2352},
	{6.5579, 9.3895, 12.9896},
	{6.7735, 9.6051, 14.7440},
	{8.5279, 11.3595, 16.4984},
	{8.7435, 11.5751, 18.2528},
	{8.9591, 13.3295, 20.0072},
	{10.7135, 13.5451, 21.7616},
	{10.9291, 15.2995, 23.5160},
	{11.1447, 15.5151, 25.2704},
	{12.8991, 17.2695, 27.0248},
	{13.1147, 17.4851, 28.7792},
	{13.3303, 19.2395, 30.5336},
	{15.0847, 19.4551, 32.2880},
	{15.3003, 21.2095, 34.0424},
	{15.5159, 21.4251, 35.7968},
	{17.2703, 23.1795, 37.5512},
	{17.4859, 23.3951, 39.3056},
	{17.7015, 25.1495, 41.0600},
	{19.4559, 25.3651, 42.8144},
	{19.6715, 27.1195, 44.5688},
	{19.8871, 27.3351, 46.3232},
	{21.6415, 29.0895, 48.0776},
	{21.8571, 29.3051, 49.8320},
	{22.0727, 31.0595, 51.5864},
	{23.8271, 31.2751, 53.3408},
	{24.0427, 33.0295, 55.0952},
	{24.2583, 33.2451, 56.8496},
	{26.0127, 34.9995, 58.6040},
	{26.2283, 35.2151, 60.3584},
	{26.4439, 36.9695, 62.1128},
	{28.1983, 37.1851, 63.8672},
	{28.4139, 38.9395, 65.6216},
	{28.6295, 39.1551, 67.3760},
	{30.3839, 40.9095, 69.1304},
	{30.5995, 41.1251, 70.8848},
	{30.8151, 42.8795, 72.6392},
	{32.5695, 43.0951, 74.3936},
	{32.7851, 44.8495, 76.1480},
	{33.0007, 45.0651, 77.9024},
	{34.7551, 46.8195, 79.6568},
	{34.9707, 47.0351, 81.4112},
	{35.1863, 48.7895, 83.1656},
	{36.9407, 49.0051, 84.9200},
	{37.1563, 50.7595, 86.6744},
	{37.3719, 50.9751, 88.4288},
	{39.1263, 52.7295, 90.1832},
	{39.3419, 52.9451, 91.9376},
	{39.5575, 54.6995, 93.6920},
	{41.3119, 54.9151, 95.4464},
	{41.5275, 56.6695, 97.2008},
	{41.7431, 56.8851, 98.9552},
	{43.4975, 58.6395, 100.7096},
	{43.7131, 58.8551, 102.4640},
	{43.9287, 60.6095, 104.2184},
	{45.6831, 60.8251, 105.9728},
	{45.8987, 62.5795, 107.7272},
	{46.1143, 62.7951, 109.4816},
	{47.8687, 64.5495, 111.2360},
	{48.0843, 64.7651, 112.9904},
	{48.2999, 66.5195, 114.7448},
	{50.0543, 66.7351, 116.4992},
	{50.2699, 68.4895, 118.2536},
	{50.4855, 68.7051, 120.0080},
	{52.2399, 70.4595, 121.7624},
	{52.4555, 70.6751, 123.5168},
	{52.6711, 72.4295, 125.2712},
	{54.4255, 72.6451, 127.0256},
	{54.6411, 74.3995, 128.7800},
	{54.8567, 74.6151, 130.5344},
	{56.6111, 76.3695, 132.2888},
	{56.8267, 76.5851, 134.0432},
	{57.0423, 78.3395, 135.7976},
	{58.7967, 78.5551, 137.5520},
	{59.0123, 80.3095, 139.3064},
	{59.2279, 80.5251, 141.0608},
	{60.9823, 82.2795, 142.8152},
	{61.1979, 82.4951, 144.5696},
	{61.4135, 84.2495, 146.3240},
	{63.1679, 84.4651, 148.0784},
	{63.3835, 86.2195, 149.8328},
	{63.5991, 86.4351, 151.5872},
	{65.3535, 88.1895, 153.3416},
	{65.5691, 88.4051, 155.0960},
	{65.7847, 90.1595, 156.8504},
	{67.5391, 90.3751, 158.6048},
	{67.7547, 92.1295, 160.3592},
	{67.9703, 92.3451, 162.1136},
	{69.7247, 94.0995, 163.8680},
	{69.9403, 94.3151, 165.6224},
	{70.1559, 96.0695, 167.3768},
	{71.9103, 96.2851, 169.1312},
	{72.1259, 98.0395, 170.8856},
}

var MEDIUM_PARCEL_TABLE = []float64{

	0.0000,
	0.7357,
	0.9304,
	1.6610,
	1.8606,
	3.0778,
	3.2934,
	4.5106,
	4.7262,
	5.9434,
	6.1590,
	7.3762,
	7.5918,
	8.8090,
	9.0246,
	10.2418,
	10.4574,
	11.6746,
	11.8902,
	13.1074,
	13.3230,
	14.5402,
	14.7558,
	15.9730,
	16.1886,
	17.4058,
	17.6214,
	18.8386,
	19.0542,
	20.2714,
	20.4870,
	21.7042,
	21.9198,
	23.1370,
	23.3526,
	24.5698,
	24.7854,
	26.0026,
	26.2182,
	27.4354,
	27.6510,
	28.8682,
	29.0838,
	30.3010,
	30.5166,
	31.7338,
	31.9494,
	33.1666,
	33.3822,
	34.5994,
	34.8150,
	36.0322,
	36.2478,
	37.4650,
	37.6806,
	38.8978,
	39.1134,
	40.3306,
	40.5462,
	41.7634,
	41.9790,
	43.1962,
	43.4118,
	44.6290,
	44.8446,
	46.0618,
	46.2774,
	47.4946,
	47.7102,
	48.9274,
	49.1430,
	50.3602,
	50.5758,
	51.7930,
	52.0086,
	53.2258,
	53.4414,
	54.6586,
	54.8742,
	56.0914,
	56.3070,
	57.5242,
	57.7398,
	58.9570,
	59.1726,
	60.3898,
	60.6054,
	61.8226,
	62.0382,
	63.2554,
	63.4710,
	64.6882,
	64.9038,
	66.1210,
	66.3366,
	67.5538,
	67.7694,
	68.9866,
	69.2022,
	70.4194,
}

/********************** END Reload Time Standards **************/

/********************** StrapOut Time Standards **************/

const RUBBER_BAND_TIME = 0.099
const RUBBER_BAND_PIECES = 32.0

const TRAY_TIME = 0.2458
const TRAY_PIECES = 220

const HANDFUL_TIME = 0.0291
const HANDFUL_PIECES = 32

const REACH_TIME = 0.0326

/********************** USPS Vehicle Time Standards **************/

//RRECS Time Standards
const INSPECT_GOV_VEHICLE_TIME = 1.8716

const FUEL_VEHICLE_MILES = 100.0
const FUEL_VEHICLE_TIME = 3.1421

/********************** Walk Time Standards **************/

const AVG_FEET_PER_STEP = 2.5    //this is normal height normal gate
const WALKING_STANDARD = 0.00429 //minute per foot

/******************** PFD Time Standards **********************/

// 3.23% added to actual time events, except the end of shift
const PERSONAL_TIME_PERCENT = 0.0323

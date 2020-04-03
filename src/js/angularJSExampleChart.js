var ngModule = angular.module('app');

ngModule.controller('angularJSExampleChartCtrl', ['$scope', '$http', '$q', '$filter', function ($scope, $http, $q, $filter) {

    $scope.readVisitors = readVisitors;
    $scope.selectIPLocation = selectIPLocation;

    $scope.visitors = [];
    $scope.monthlyVisitors = [];
    $scope.currentMonthTotal = undefined;
    $scope.currentPageMonthTotal = undefined;
    $scope.currentMonthName = "";
    $scope.chartData = [];
    $scope.currentMonth = "";
    $scope.ipLocationList = [];
    $scope.selectedIP = [];

    $scope.pingTime = 0.0;
    $scope.distance = 0.0;
    $scope.browserCity = "";
    $scope.browserState = "";

    function ping() {
        let start = performance.now();
        $http.get("/ping").then(function (res) {
            let finish = performance.now();
            $scope.pingTime = Math.round(finish - start);
        }, function(error) {
            console.log(error.data);
        });
    }

    function readIP() {
        $http.get("/readIP").then(function (res) {
            let results = res.data;
            console.log("RESULTS: ", results);
            $scope.distance = calculateDistance(results.latitude, 32.779167 , results.longitude, -96.808891);
            $scope.browserCity = results.city;
            $scope.browserState = results.region_code;
        }, function (error) {
            console.log(error);
        })
    }

    // Found this here:
    // https://stackoverflow.com/questions/27928/calculate-distance-between-two-latitude-longitude-points-haversine-formula
    function calculateDistance(lat1, lat2, long1, long2) {
        var p = 0.017453292519943295;    // Math.PI / 180
        var c = Math.cos;
        var a = 0.5 - c((lat2 - lat1) * p)/2 +
            c(lat1 * p) * c(lat2 * p) *
            (1 - c((long2 - long1) * p))/2;

        return 12742 * Math.asin(Math.sqrt(a)); // 2 * R; R = 6371 km
    }

    function drawChart(data) {
        let chart;

        chart = c3.generate({
            bindto: 'div#chart',
            size: {
                height: 300
            },
            padding: {
                top: 20,
                right: 50,
                bottom: 20,
                left: 50,
            },
            data: {
                url: '/readMonthlyVisitors',
                mimeType: 'json',
                x: 'X',
                xFormat: '%Y-%m-%dT%H:%M:%SZ',
                keys: {
                    x: 'date_stamp', // it's possible to specify 'x' when category axis
                    value: ['count', 'pageCount', 'avgCount', 'avgPageCount'],
                },
                names: {
                    count: 'Unique Visitors',
                    pageCount: 'Page Views',
                    avgCount: '3 Month Rolling AVG Unique Visitors',
                    avgPageCount: '3 Month Rolling AVG Page Views'
                },
                colors: {
                    count: '#ff0000',
                    pageCount: '#0000ff',
                    avgCount: '#ffcccb',
                    avgPageCount: '#add8e6'
                },
            },
            axis: {
                x: {
                    type: 'timeseries',
                    tick: {
                        rotate: 75,
                        multiline: false,
                        culling: false,
                        format: '%Y-%m'
                    }
                }
            },
        });
    }

    function sumObjectProperty(type) {
        let sum = 0;
        for(const[index,item] of $scope.monthlyVisitors.entries()) {
            sum += item[type];
        }
        return sum;
    }

    function getMonthName() {
        const monthNames = ["January", "February", "March", "April", "May", "June",
            "July", "August", "September", "October", "November", "December"
        ];
        const d = new Date();
        return monthNames[d.getMonth()];
    }

    function readVisitors() {
        $http.get("/readVisitors").then(function (res) {
            let results;
            results = res.data;
            $scope.visitors = results;
            //getIPLocations($scope.visitors);
            //drawChart(results);
        }, function (err) {
            alert("ERROR, /readVisitors: ", err);
        })
    }

    function readMonthlyVisitors() {
        $http.get("/readMonthlyVisitors").then(function (res) {
            let results;
            let currentMonth = getMonthName();
            results = res.data;
            $scope.monthlyVisitors = results;
            $scope.currentMonthName = currentMonth;

            // Setting the total for this month!
            for (const[index,item] of $scope.monthlyVisitors.entries()) {
                if (item.month === currentMonth) {
                    $scope.currentMonthTotal = item.count;
                    $scope.currentPageMonthTotal = item.pageCount;
                }
            }
            //$scope.currentMonthTotal = 0;
            // sort our results by month
            results.sort(function(a,b) {return (a.id > b.id) ? 1 : ((b.id > a.id) ? -1 : 0);} );
            drawChart(results);
        }, function (err) {
            alert("ERROR /readMonthlyVisitors: ", err);
        })
    }

    function setCurrentMonth() {
        const monthNames = ["January", "February", "March", "April", "May", "June",
            "July", "August", "September", "October", "November", "December"
        ];
        const d = new Date();
        $scope.currentMonth = monthNames[d.getMonth()];
    }

    function getIPLocations(ips) {
        //let ip;
        $http.get("/getIPLocation").then(function (res) {
            let results;
            results = res.data;
            $scope.ipLocationList = results;
        }, function (err) {
            alert("Error obtaining the location for that IP: ", err);
        });
    }

    function selectIPLocation(ip) {
        $scope.selectedIP = $scope.ipLocationList.filter(function (ipNumber) {
            return ipNumber.ip === ip;
        });
    }

    getIPLocations();
    setCurrentMonth();
    readVisitors();
    readMonthlyVisitors();

    readIP();
    ping();
    setInterval(ping, 2500);
}]);

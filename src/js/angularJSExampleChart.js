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

    function drawChart(data) {
        let chart;
        let countSum = sumObjectProperty("count");
        let pageCountSum = sumObjectProperty("pageCount");
        let countAvg = countSum/$scope.monthlyVisitors.length;
        let pageCountAvg = pageCountSum/$scope.monthlyVisitors.length;

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
                    value: ['count', 'pageCount'],
                },
                names: {
                    count: 'Unique Visitors',
                    pageCount: 'Page Views'
                }
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
            grid: {
                y: {
                    lines: [
                        {value: countAvg, text: 'Average Unique Visitors', position: 'middle'},
                        {value: pageCountAvg, text: 'Average Page Views', position: 'middle'}
                    ]
                }
            }
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

            console.log("IPs: ", $scope.visitors);

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
            console.log("Monthly Visitors: ", $scope.monthlyVisitors);
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
            console.log("IP LOCATION: ", results);
        }, function (err) {
            alert("Error obtaining the location for that IP: ", err);
        });
    }

    function selectIPLocation(ip) {
        $scope.selectedIP = $scope.ipLocationList.filter(function (ipNumber) {
            return ipNumber.ip === ip;
        });

        console.log("SELECTED IP: ", $scope.selectedIP);
    }

    getIPLocations();
    setCurrentMonth();
    readVisitors();
    readMonthlyVisitors();
}]);

var ngModule = angular.module('app');

ngModule.controller('angularJSExampleChartCtrl', ['$scope', '$http', '$q', '$filter', function ($scope, $http, $q, $filter) {

    $scope.readVisitors = readVisitors;
    $scope.selectIPLocation = selectIPLocation;

    $scope.visitors = [];
    $scope.monthlyVisitors = [];
    $scope.currentMonthTotal = 0;
    $scope.currentMonthName = "";
    $scope.chartData = [];
    $scope.currentMonth = "";
    $scope.ipLocationList = [];
    $scope.selectedIP = [];

    function drawChart(data) {
        let chart;

        let monthCounts;
        let times;
        let months;

        $scope.chartData = data;

        // monthCounts = getMonthCounts($scope.chartData);
        // months = $scope.chartData.map(getMonth);
        // times = $scope.chartData.map(getTimes);

        chart = c3.generate({
            bindto: 'div#chart',
            size: {
                height: 300
            },
            data: {
                x: 'x',
                xFormat: '%Y-%m',
                //        xFormat: '%Y%m%d', // 'xFormat' can be used as custom format of 'x'
                columns: [
                    ['x', '2019-01', '2019-02', '2019-03', '2019-04', '2019-05', '2019-06', '2019-07', '2019-08', '2019-09', '2019-10', '2019-11', '2019-12'],
                    ['Unique Visitors',
                        $scope.chartData[0]['count'],
                        $scope.chartData[1]['count'],
                        $scope.chartData[2]['count'],
                        $scope.chartData[3]['count'],
                        $scope.chartData[4]['count'],
                        $scope.chartData[5]['count'],
                        0,
                        0,
                        0,
                        0,
                        0,
                        0]
                ]
            },
            axis: {
                x: {
                    type: 'category',
                    tick: {
                        culling: false
                        //format: '%Y-%m-%d'
                    }
                }
            }
        });

        chart.load($scope.visitors);
    }

    function getTimes(obj) {
        return Date.parse(obj.timestamp);
    }

    function getMonthName() {
        const monthNames = ["January", "February", "March", "April", "May", "June",
            "July", "August", "September", "October", "November", "December"
        ];
        const d = new Date();
        return monthNames[d.getMonth()];
    }

    function getMonth(obj) {
        let month;
        let date;

        date = new Date(obj.timestamp);

        month = date.getMonth();
        return date;
    }

    function getMonthCounts(obj) {
        let monthCount = [0,0,0,0,0,0,0,0,0,0,0,0];
        let date;

        for(const[index,item] of obj.entries()) {
            date = new Date(item.timestamp);
            monthCount[date.getMonth()] += 1;
        }

        console.log("DATE COUNT: ", monthCount);
        return monthCount;
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
                }
            }
            //$scope.currentMonthTotal = 0;
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
    // drawChart();
}]);

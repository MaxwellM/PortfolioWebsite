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
<<<<<<< Updated upstream

        let monthCounts;
        let times;
        let months;

        // Lets fill out our array with 0s, then fill it with data!
        let n = 12;
        for (var i=0; i < n; i++) {
            $scope.chartData.push({count: 0, pageCount: 0});
        }

        for (const [index,item] of data.entries()) {
            $scope.chartData[index] = item;
        }

        //$scope.chartData.fill(0, 0, 11);
        //$scope.chartData.push(data);

        // monthCounts = getMonthCounts($scope.chartData);
        // months = $scope.chartData.map(getMonth);
        // times = $scope.chartData.map(getTimes);
=======
        let countSum = sumObjectProperty("count");
        let pageCountSum = sumObjectProperty("pageCount");
        let countAvg = countSum/$scope.monthlyVisitors.length;
        let pageCountAvg = pageCountSum/$scope.monthlyVisitors.length;
>>>>>>> Stashed changes

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
<<<<<<< Updated upstream
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
                        $scope.chartData[6]['count'],
                        $scope.chartData[7]['count'],
                        $scope.chartData[8]['count'],
                        $scope.chartData[9]['count'],
                        $scope.chartData[10]['count'],
                        $scope.chartData[11]['count']],
                    ['Page Views',
                        $scope.chartData[0]['pageCount'],
                        $scope.chartData[1]['pageCount'],
                        $scope.chartData[2]['pageCount'],
                        $scope.chartData[3]['pageCount'],
                        $scope.chartData[4]['pageCount'],
                        $scope.chartData[5]['pageCount'],
                        $scope.chartData[6]['pageCount'],
                        $scope.chartData[7]['pageCount'],
                        $scope.chartData[8]['pageCount'],
                        $scope.chartData[9]['pageCount'],
                        $scope.chartData[10]['pageCount'],
                        $scope.chartData[11]['pageCount']]
                ]
=======
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
>>>>>>> Stashed changes
            },
            axis: {
                x: {
                    type: 'category',
                    tick: {
<<<<<<< Updated upstream
                        culling: false
                        //format: '%Y-%m-%d'
=======
                        rotate: 75,
                        multiline: false,
                        culling: false,
                        format: '%Y-%m'
>>>>>>> Stashed changes
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
<<<<<<< Updated upstream

        chart.load($scope.visitors);
=======
>>>>>>> Stashed changes
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

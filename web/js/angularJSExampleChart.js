var ngModule = angular.module('app');

ngModule.controller('angularJSExampleChartCtrl', ['$scope', '$http', '$q', '$filter', function ($scope, $http, $q, $filter) {

    $scope.readVisitors = readVisitors;

    $scope.visitors = [];
    $scope.monthlyVisitors = [];
    $scope.chartData = [];
    $scope.currentMonth = "";

    function drawChart(data) {
        let chart;

        let monthCounts;
        let times;
        let months;

        $scope.chartData = data;

        monthCounts = getMonthCounts($scope.chartData);
        months = $scope.chartData.map(getMonth);
        times = $scope.chartData.map(getTimes);

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
                        monthCounts[0],
                        monthCounts[1],
                        monthCounts[2],
                        monthCounts[3],
                        monthCounts[4],
                        monthCounts[5],
                        monthCounts[6],
                        monthCounts[7],
                        monthCounts[8],
                        monthCounts[9],
                        monthCounts[10],
                        monthCounts[11]]
                ]
            },
            axis: {
                x: {
                    type: 'category',
                    tick: {
                        culling:  {
                            max: 12
                        },
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

    function getMonth(obj) {
        const monthNames = ["January", "February", "March", "April", "May", "June",
            "July", "August", "September", "October", "November", "December"
        ];
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

    // function getVisitCount(data) {
    //     const monthNames = ["January", "February", "March", "April", "May", "June",
    //         "July", "August", "September", "October", "November", "December"
    //     ];
    //     let monthCounts = [];
    //     let date;
    //     let month;
    //
    //     for(const[index,item] of data.entries()) {
    //         date = new Date(item.timestamp);
    //         month = date.getMonth();
    //     }
    //
    //     return monthCounts;
    // }

    // function getMonths(data) {
    //     const monthNames = ["January", "February", "March", "April", "May", "June",
    //         "July", "August", "September", "October", "November", "December"
    //     ];
    //     let months = [];
    //     let date;
    //     let month;
    //
    //     console.log("DATA: ", data);
    //
    //     for(const[index,item] of data.entries()) {
    //         date = new Date(item.timestamp);
    //         month = date.getMonth();
    //         if (!months.includes(monthNames[month])) {
    //             months.push(monthNames[month]);
    //         } else {
    //             console.log("Already have that month: ", month, months)
    //         }
    //     }
    //     //let month = data.getMonth();
    //
    //     console.log("MONTHS: ", months);
    //
    //     return months;
    // }

    function readVisitors() {
        $http.get("/readVisitors").then(function (res) {
            let results;
            results = res.data;
            $scope.visitors = results;

            console.log("IPs: ", $scope.visitors);

            drawChart(results);
        }, function (err) {
            alert("ERROR, /readVisitors: ", err);
        })
    }

    function readMonthlyVisitors() {
        $http.get("/readMonthlyVisitors").then(function (res) {
            let results;
            results = res.data;
            $scope.monthlyVisitors = results;

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

    setCurrentMonth();
    readVisitors();
    readMonthlyVisitors();
    // drawChart();
}]);

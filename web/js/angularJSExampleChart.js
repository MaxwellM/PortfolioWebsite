var ngModule = angular.module('app');

ngModule.controller('angularJSExampleChartCtrl', ['$scope', '$http', '$q', '$filter', function ($scope, $http, $q, $filter) {

    $scope.readVisitors = readVisitors;

    $scope.visitors = [];

    function drawChart() {
        let months = getMonths($scope.visitors);

        var chart = c3.generate({
            bindto: '#chart',
            data: {
                columns: [
                    ['data1', 30, 200, 100, 400, 150, 250],
                    ['data2', 50, 20, 10, 40, 15, 25]
                ]
            }
        });
    }

    function getMonths(data) {
        const monthNames = ["January", "February", "March", "April", "May", "June",
            "July", "August", "September", "October", "November", "December"
        ];
        let months = [];
        let date;
        let month;

        console.log("DATA: ", data);

        for(const[index,item] of data.entries()) {
            date = new Date(item.timestamp);
            month = date.getMonth();
            if (!months.includes(monthNames[month])) {
                months.push(monthNames[month]);
            } else {
                console.log("Already have that month: ", month, months)
            }
        }
        //let month = data.getMonth();

        console.log("MONTHS: ", months);

        return months;
    }

    function readVisitors() {
        $http.get("/readVisitors").then(function (res) {
            let results;
            results = res.data;
            $scope.visitors = results;

            console.log("IPs: ", $scope.visitors);

            drawChart();
        }, function (err) {
            alert("ERROR, /readVisitors: ", err);
        })
    }

    readVisitors();
    // drawChart();
}]);

var ngModule = angular.module('app');

ngModule.controller('angularJSExampleChartCtrl', ['$scope', '$http', '$q', '$filter', function ($scope, $http, $q, $filter) {

    $scope.readVisitors = readVisitors;

    $scope.allVisitors = [];
    $scope.visitors = [];

    console.log("VISITOR LOADED!");

    // function retrieveVisitors() {
    //     $http.get("/getVisitors").then(function (res) {
    //         let results;
    //         results = res.data;
    //         $scope.allVisitors = results;
    //         console.log("VISITORS: ", results);
    //     }, function(error) {
    //         alert(error.data);
    //     });
    // }

    function drawChart() {
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

    function visitorCounter() {
        $http.get("/visitorCounter").then(function (res) {
        }, function (err) {
            alert("ERROR /visitorCounter: ", err);
        })
    }

    function readVisitors() {
        $http.get("/readVisitors").then(function (res) {
            let results;
            results = res.data;
            $scope.visitors = results;

            console.log("IPs: ", $scope.visitors);
        }, function (err) {
            alert("ERROR, /readVisitors: ", err);
        })
    }

    visitorCounter();
    readVisitors();
    drawChart();
}]);

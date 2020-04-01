var ngModule = angular.module('app');

ngModule.controller('footerCtrl', ['$scope', '$http', '$q', '$filter', function ($scope, $http, $q, $filter) {

    $scope.readVisitors = readVisitors;

    $scope.visitors = [];
    $scope.pingTime = 0.0;
    $scope.distance = 0.0;
    $scope.browserCity = "";
    $scope.browserState = "";

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
        }, function (err) {
            alert("ERROR, /readVisitors: ", err);
        })
    }

    function ping() {
        let start = performance.now();
        $http.get("/ping").then(function (res) {
            let finish = performance.now();
            let results;
            results = res.data;
            $scope.pingTime = (finish - start);
        }, function(error) {
            alert(error.data);
        });
    }

    function readIP() {
        $http.get("/readIP").then(function (res) {
            let results = res.data;
            console.log("RESULTS: ", results);
            //$scope.distance = calculateDistance(results.latitude, 32.779167 , results.longitude, -96.808891);
            //$scope.browserCity = results.city;
            //$scope.browserState = results.region_code;
        }, function (error) {
            console.log(error);
        })
    }

    // Found this here:
    // https://stackoverflow.com/questions/27928/calculate-distance-between-two-latitude-longitude-points-haversine-formula
    // function calculateDistance(lat1, lat2, long1, long2) {
    //     var p = 0.017453292519943295;    // Math.PI / 180
    //     var c = Math.cos;
    //     var a = 0.5 - c((lat2 - lat1) * p)/2 +
    //         c(lat1 * p) * c(lat2 * p) *
    //         (1 - c((long2 - long1) * p))/2;
    //
    //     return 12742 * Math.asin(Math.sqrt(a)); // 2 * R; R = 6371 km
    // }

    readIP();
    visitorCounter();
    readVisitors();
    ping();
    setInterval(ping, 2500);
}]);

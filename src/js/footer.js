var ngModule = angular.module('app');

ngModule.controller('footerCtrl', ['$scope', '$http', '$q', '$filter', function ($scope, $http, $q, $filter) {

    $scope.readVisitors = readVisitors;

    $scope.visitors = [];

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

    visitorCounter();
    readVisitors();
}]);

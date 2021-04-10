var ngModule = angular.module('app');

ngModule.controller('footerCtrl', ['$scope', '$http', '$q', '$filter', function ($scope, $http, $q, $filter) {

    $scope.visitors = [];

    function getVisitorInfo() {
        $http.get("/api/visitorCounter").then(function (res) {
            $http.get("/api/readVisitors").then(function (res) {
                let results;
                results = res.data;
                $scope.visitors = results;
            }, function (err) {
                alert("ERROR, /readVisitors: ", err);
            })
        }, function (err) {
            alert("ERROR /visitorCounter: ", err);
        })
    }

    getVisitorInfo();
}]);

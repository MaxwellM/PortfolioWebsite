var ngModule = angular.module('app');

ngModule.controller('angularJSExampleDisplayCtrl', ['$scope', '$http', '$q', '$filter', function ($scope, $http, $q, $filter) {

    $scope.getWeather = getWeather;

    // $scope.weatherData = [];
    // $scope.weatherLocation = "";

    $scope.weatherRequestCount;

    function getWeather(location) {
        // We're giving out request to the backend so that we can monitor how many requests are made within a day.
        // We're cheap and are on the free plan which means we have 50 requests a day. We'll implement a count in Go
        // that will cease requests if they exceed that amount.

        $http.get("/getWeather", {params:{location}}).then(function (res) {
            let results;
            results = res.data;
            $scope.weatherData = results;
            console.log("BACKEND RESULTS: ", results);
        }, function(error) {
           alert(error.data);
        });
    }


}]);

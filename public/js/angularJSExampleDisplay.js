var ngModule = angular.module('app');

ngModule.controller('angularJSExampleDisplayCtrl', ['$scope', '$http', '$q', '$filter', function ($scope, $http, $q, $filter) {

    $scope.getWeather = getWeather;
    $scope.trimDate = trimDate;
    $scope.getDayName = getDayName;

    $scope.weatherData = [];
    // $scope.weatherLocation = "";

    $scope.weatherRequestCount;
    $scope.currentConditionsWeather;

    function getWeather(location) {
        // We're giving out request to the backend so that we can monitor how many requests are made within a day.
        // We're cheap and are on the free plan which means we have 50 requests a day. We'll implement a count in Go
        // that will cease requests if they exceed that amount.

        $http.get("/getWeather", {params:{location}}).then(function (res) {
            let results;
            results = res.data;
            $scope.weatherData = results.Forecast;
            console.log("FORECAST RESULTS: ", $scope.weatherData);
            console.log("DATE: ", $scope.weatherData[0].Date);

            $scope.currentConditionsWeather = results.Current;
            console.log("CURRENT RESULTS: ", $scope.currentConditionsWeather);
            readLocalWeatherReport();
        }, function(error) {
           alert(error.data);
        });
    }

    function readLocalWeatherReport() {
        $http.get("/getLocalWeather", {params:{location}}).then(function (res) {
            let results;
            results = res.data;
            $scope.weatherData = results;
            console.log("BACKEND RESULTS: ", results);
        }, function(error) {
            alert(error.data);
        });
    }

    function readCurrentConditions() {
        $http.get("/getLocalCurrentConditions", {params:{location}}).then(function (res) {
            let results;
            results = res.data;
            $scope.currentConditionsWeather = results.Result[0];
            console.log("CURRENT CONDITIONS: ", results);
        }, function(error) {
            alert(error.data);
        });
    }

    function trimDate(date) {
        return date.substring(0, date.indexOf("T"));
    }

    function getDayName(date) {
        date = date.substring(0, date.indexOf("T"));
        let Datedate = new Date(date);
        //date = date.substring(0, date.indexOf("T"));
        let days = ['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday'];
        //console.log("DAY: ", Datedate.getDay());
        let newDate = days[Datedate.getDay()];
        return newDate;
    }

    readLocalWeatherReport();
    readCurrentConditions();

}]);
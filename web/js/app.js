var ngModule = angular.module('app', ['720kb.datepicker']);

ngModule.controller('myCtrl', ['$scope', '$http', '$q', '$filter', function ($scope, $http, $q, $filter) {

    $scope.visitors = [];

    function myMap() {
        var mapOptions = {
            center: new google.maps.LatLng(40.586667, -111.861244),
            zoom: 10,
            mapTypeId: google.maps.MapTypeId.HYBRID
        };
        var map = new google.maps.Map(document.getElementById("map"), mapOptions);
    }

    function visitorCounter() {
        $http.get("/visitorCounter").then(function (res) {
            let results;
            results = res.data;
            $scope.visitors = results;

            console.log("IPs: ", $scope.visitors.length);
        }, function (err) {
            alert("ERROR: ", err);
        })
    }

    visitorCounter();
    myMap();

}]);

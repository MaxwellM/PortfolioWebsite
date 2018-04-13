var ngModule = angular.module('app', []);

ngModule.controller('myCtrl', ['$scope', '$http', '$q', '$filter', function ($scope, $http, $q, $filter) {

    function myMap() {
        var mapOptions = {
            center: new google.maps.LatLng(40.586667, -111.861244),
            zoom: 10,
            mapTypeId: google.maps.MapTypeId.HYBRID
        };
        var map = new google.maps.Map(document.getElementById("map"), mapOptions);
    }

    myMap();

}]);
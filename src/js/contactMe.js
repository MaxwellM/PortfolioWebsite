var ngModule = angular.module('app');

ngModule.controller('contactMeCtrl', ['$scope', '$http', '$q', '$filter', function ($scope, $http, $q, $filter) {

    $scope.sendMessage = sendMessage;

    $scope.user = {
        email: '',
        phone: '',
        firstName: '',
        lastName: '',
        message: '',
    };

    $scope.emailRegex = /^.+@.+\..+$/;
    $scope.phoneRegex = /^[(][0-9]{3}[)] [0-9]{3}-[0-9]{4}$/;

    function sendMessage() {
        let obj = {
            name: $scope.user.firstName + " " + $scope.user.lastName,
            email: $scope.user.email,
            phone: $scope.user.phone,
            message: $scope.user.message
        };

        // Checking to make sure all fields are filled out!
        if (!checkObject()) {
            alert("Missing Field!");
            return
        }

        $http.post('/sendMessage', obj).then(function (res) {
            let results = res.data;
            alert("Message Successfully Sent!");
        }, function (err) {
            console.log("Error sending message: ", err)
        })
    }

    function checkObject() {
        let result = !Object.values($scope.user).every(o => o === "");
        return result;
    }

    function myMap() {
        let mapOptions = {
            center: new google.maps.LatLng(40.586667, -111.861244),
            zoom: 10,
            mapTypeId: google.maps.MapTypeId.HYBRID
        };
        let map = new google.maps.Map(document.getElementById("map"), mapOptions);
    }

    myMap();


}]);

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
    $scope.phoneRegex = /^[\+]?[(]?[0-9]{3}[)]?[-\s\.]?[0-9]{3}[-\s\.]?[0-9]{4,6}$/im;


    // IF YOU CANNOT SEND THE MAIL, JUST VERIFY IN THE GOOGLE ACCOUNT THAT LESS SECURE APPS ARE ENABLED
    // https://stackoverflow.com/questions/10013736/how-can-i-avoid-google-mail-server-asking-me-to-log-in-via-browser
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

        $http.post('/api/sendMessage', obj).then(function (res) {
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
            center: new google.maps.LatLng(41.233219, -112.027877),
            zoom: 10,
            mapTypeId: google.maps.MapTypeId.HYBRID
        };
        let map = new google.maps.Map(document.getElementById("map"), mapOptions);
    }

    myMap();


}]);

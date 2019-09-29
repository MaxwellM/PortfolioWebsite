var ngModule = angular.module('app');

ngModule.controller('contactMeCtrl', ['$scope', '$http', '$q', '$filter', function ($scope, $http, $q, $filter) {

    $scope.sendMessage = sendMessage;

    // $scope.name = "Name";
    // $scope.email = "Email";
    // $scope.phone = "Phone Number";
    // $scope.message = "Message";

    function sendMessage(name, email, phone, message) {
        let obj = {
            name: name,
            email: email,
            phone: phone,
            message: message
        };

        // Checking to make sure all fields are filled out!
        if (!checkArguments(arguments)) {
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

    // This will check that the email looks like an email. Found this here:
    // https://stackoverflow.com/questions/46155/how-to-validate-an-email-address-in-javascript
    function validEmail(email) {
        let re = /^[a-zA-Z0-9\-_]+(\.[a-zA-Z0-9\-_]+)*@[a-z0-9]+(\-[a-z0-9]+)*(\.[a-z0-9]+(\-[a-z0-9]+)*)*\.[a-z]{2,4}$/;
        return re.test(email);
    }

    function checkArguments(arguments) {
        for (const property in arguments) {
            console.log("THIS: ", arguments[property]);
            if (arguments[property] === undefined) {
                return false;
            } else if (parseInt(property) === 1) {
                // this is the email one
                if (!validEmail(arguments[property])) {
                    return false;
                }
            }
        }
        return true;
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

var ngModule = angular.module('app');

ngModule.controller('contactMeCtrl', ['$scope', '$http', '$q', '$filter', function ($scope, $http, $q, $filter) {

    $scope.sendMessage = sendMessage;

    $scope.name = "Name";
    $scope.email = "Email";
    $scope.phone = "Phone Number";
    $scope.message = "Message";

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

    function checkArguments(arguments) {
        for (const [index,item] of arguments) {
            if (item === undefined) {
                return false;
            }
        }
        return true;

        // for (var key in obj) {
        //     console.log(obj[key]);
        // }
    }

}]);

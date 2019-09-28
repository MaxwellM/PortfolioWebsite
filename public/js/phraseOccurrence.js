var ngModule = angular.module('app');

ngModule.controller('goExampleOccurrenceCtl', ['$scope', '$http', '$q', '$filter', function ($scope, $http, $q, $filter) {

    $scope.findOccurrences = findOccurrences;
    $scope.occurrences = "";

    function findOccurrences(str) {
        // Splitting string
        let SplitString = str.split(" ");

        $http.get("/getOccurrences", {params:{SplitString}}).then(function (res) {
            let results;
            results = res.data;
            $scope.occurrences = JSON.stringify(results, null, 2);

        }, function(error) {
            alert(error.data);
        });
    }

}]);

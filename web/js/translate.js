var ngModule = angular.module('app');

ngModule.controller('goExampleTranslateCtl', ['$scope', '$http', '$q', '$filter', '$sanitize', function ($scope, $http, $q, $filter, $sanitize) {

    $scope.translate = translate;
    $scope.translation = "";

    function translate(str) {
        // Splitting string
        let SplitString = str;

        $http.get("/translate", {params:{SplitString}}).then(function (res) {
            let results;
            results = res.data;
            $scope.translation = results;

        }, function(error) {
            alert(error.data);
        });
    }

}]);

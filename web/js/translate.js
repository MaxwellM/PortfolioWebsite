var ngModule = angular.module('app');

ngModule.controller('goExampleTranslateCtl', ['$scope', '$http', '$q', '$filter', '$sanitize', function ($scope, $http, $q, $filter, $sanitize) {

    $scope.translate = translate;
    $scope.translation = "";
    $scope.languages = [];

    function translate(str, lg) {
        // Splitting string
        let SplitString = str;
        let Lang = lg;

        $http.get("/translate", {params:{SplitString, Lang}}).then(function (res) {
            let results;
            results = res.data;
            $scope.translation = results;

        }, function(error) {
            alert(error.data);
        });
    }

    function fillLanguages() {
        $scope.languages = [
            {Name:"French", code: "fr"},
            {Name:"Esperanto", code: "eo"},
            {Name:"Arabic", code: "ar"},
            {Name:"Chinese", code: "zh"},
            {Name:"Russian", code:"ru"},
            {Name:"German", code: "de"},
            {Name:"Latin", code: "la"},
            {Name:"Vietnamese", code: "vi"},
        ];
        //$scope.languages.code = $scope.languages[0];
    }

    fillLanguages();

}]);

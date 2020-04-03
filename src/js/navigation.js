var ngModule = angular.module('app');

ngModule.controller('navigationCtrl', ['$scope', '$http', '$q', '$filter', '$window', '$location', function ($scope, $http, $q, $filter, $window, $location) {

    //$scope.currentNavItem = '';
    //$scope.currentMenuItem = '';

    $scope.examples = [
        "Table Example",
        "Display Example",
        "Chart Example",
        "String Duplication",
        "Translate",
        "Twitter",
        "Unity Lost In Space"
    ];

}]);

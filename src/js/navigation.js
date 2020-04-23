var ngModule = angular.module('app');

ngModule.controller('navigationCtrl', ['$scope', '$rootScope', '$http', '$q', '$filter', function ($scope, $rootScope, $http, $q, $filter) {
    // This is our default page, the main page!
    $rootScope.buttonName = "Maxwell Morin";
    $scope.currentNavItem = "Maxwell Morin";
    $rootScope.currentMenuItem = "";

    $scope.examples = [
        "Table Example",
        "Display Example",
        "Chart Example",
        "String Duplication",
        "Translate",
        "Twitter",
        "Unity Lost In Space",
        "Unity Eggcite Bike",
        "Unity Eggman",
        "Stock Tracker"
    ];
}]);

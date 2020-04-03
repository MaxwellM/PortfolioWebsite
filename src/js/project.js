var ngModule = angular.module('app');

ngModule.controller('projectCtrl', ['$scope', '$http', '$q', '$filter', function ($scope, $http, $q, $filter) {

    //$scope.currentNavItem = $scope.$parent.currentNavItem;
    $scope.$parent.currentNavItem = "Projects";
    //navigationCtrl.currentNavItem = "Projects";
}]);

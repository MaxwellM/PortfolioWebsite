var ngModule = angular.module('app');

ngModule.controller('mainPageCtrl', ['$scope', '$http', '$q', '$filter', function ($scope, $http, $q, $filter) {

    $scope.$parent.currentNavItem = "Maxwell Morin";
    GitHubCalendar(".calendar", "MaxwellM", { responsive: true });

}]);

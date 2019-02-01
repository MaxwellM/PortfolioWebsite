var ngModule = angular.module('app');

ngModule.controller('mainPageCtrl', ['$scope', '$http', '$q', '$filter', function ($scope, $http, $q, $filter) {
    console.log("mainPageCtrl LOADED OKAY!")

}]);

var ngModule = angular.module('app', ['720kb.tooltips', 'ngSanitize', 'ngtweet', 'ngMaterial', 'ngMessages', 'md.data.table']).config(function ($routeProvider, $locationProvider) {
    $locationProvider.html5Mode(true);
    $routeProvider.when('/', {
        templateUrl : './html/landingPages/mainLanding.html',
        controller : 'mainPageCtrl'
    }).when('/resume', {
        templateUrl : './html/landingPages/resumeLanding.html',//actual location will vary according to your local folder structure
        controller : 'resumeCtrl'
    }).when('/projects', {
        templateUrl : './html/landingPages/projectsLanding.html',
        controller : 'projectCtrl'
    });
});

ngModule.controller('myCtrl', ['$scope', '$http', '$q', '$filter', '$window', function ($scope, $http, $q, $filter, $window) {

    //$scope.currentNavItem = "Maxwell Morin";

}]);
var ngModule = angular.module('app', ['720kb.tooltips', 'ngSanitize', 'ngtweet', 'ngMaterial', 'ngMessages', 'ngCookies', 'md.data.table'])

.config(['$mdThemingProvider', function ($mdThemingProvider) {
    'use strict';

    $mdThemingProvider.theme('docs-dark', 'default')
        .primaryPalette('grey')
        .warnPalette('red')
        .accentPalette('blue')
        .dark();
}]);

ngModule.controller('myCtrl', ['$scope', '$http', '$q', '$filter', '$cookies', function ($scope, $http, $q, $filter, $cookies) {

    // Lets get our cookie on page load!
    function getCookie() {
        $http.get("/getCookie").then(function (res) {
            $cookies.put("Maxwell_Ross_Morin");
        }, function (error) {

        })
    }

    getCookie();
}]);

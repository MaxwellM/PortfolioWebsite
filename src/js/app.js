var ngModule = angular.module('app', ['720kb.tooltips', 'ngSanitize', 'ngtweet', 'ngMaterial', 'ngMessages', 'ngCookies', 'md.data.table', 'oc.lazyLoad', 'ui.router'])

.config(['$mdThemingProvider', '$stateProvider', '$locationProvider', '$urlRouterProvider', function ($mdThemingProvider, $stateProvider, $locationProvider, $urlRouterProvider) {
    'use strict';

    $mdThemingProvider.theme('docs-dark', 'default')
        .primaryPalette('grey')
        .warnPalette('red')
        .accentPalette('blue')
        .dark();

    $locationProvider.hashPrefix(''); // by default '!'
    $locationProvider.html5Mode(true);

    // creating routes or states
    $stateProvider
        .state('Home', {
            url : '/',
            templateUrl : "./html/includes/main.html",
            controller : "mainPageCtrl",
        })
        .state('ContactMe', {
            url : '/contactMe',
            templateUrl : "./html/includes/contactMePage.html",
            controller : "contactMeCtrl"
        })
        .state('Resume', {
            url : '/resume',
            templateUrl : "./html/includes/resume.html",
            controller : "resumeCtrl"
        })
        .state('Projects', {
            url : '/projects',
            templateUrl : "./html/includes/projectPage.html",
            controller : "projectCtrl"
        })
        .state('AngularJSExampleChart', {
            url : '/angularJsExampleChart',
            templateUrl : "./html/includes/angularJSExampleChart.html",
            controller : "angularJSExampleChartCtrl"
        })
        .state('AngularJSExampleDisplay', {
            url : '/angularJsExampleDisplay',
            templateUrl : "./html/includes/angularJSExampleDisplay.html",
            controller : "angularJSExampleDisplayCtrl"
        })
        .state('AngularJSExampleTable', {
            url : '/angularJsExampleTable',
            templateUrl : "./html/includes/angularJSExampleTable.html",
            controller : "angularJSExampleTableCtrl"
        })
        .state('PhraseOccurrence', {
            url : '/phraseOccurrence',
            templateUrl : "./html/includes/phraseOccurrence.html",
            controller : "goExampleOccurrenceCtl"
        })
        .state('StockTracker', {
            url : '/stockTracker',
            templateUrl : "./html/includes/stockTracker.html",
            controller : "stockTrackerCtrl"
        })
        .state('Translate', {
            url : '/translate',
            templateUrl : "./html/includes/translate.html",
            controller : "goExampleTranslateCtl"
        })
        .state('Twitter', {
            url : '/twitter',
            templateUrl : "./html/includes/twitter.html",
            controller : "goExampleTwitterCtl"
        })
        .state('UnityEggciteBike', {
            url : '/unityEggciteBike',
            templateUrl : "./html/includes/unityEggciteBike.html",
            controller : "unityEggciteBikeCtrl"
        })
        .state('UnityEggman', {
            url : '/unityEggman',
            templateUrl : "./html/includes/unityEggman.html",
            controller : "unityEggmanCtrl"
        })
        .state('UnityLostInSpace', {
            url : '/unityLostInSpace',
            templateUrl : "./html/includes/unityLostInSpace.html",
            controller : "unityLostInSpaceCtrl"
        });

        // Redirect to home page if url does not
        // matches any of the three mentioned above
        $urlRouterProvider.otherwise("/");
}]);

ngModule.controller('myCtrl', ['$scope', '$http', '$q', '$filter', '$cookies', function ($scope, $http, $q, $filter, $cookies) {

    // Lets get our cookie on page load!
    function getCookie() {
        $http.get("/api/getCookie").then(function (res) {
            $cookies.put("Maxwell_Ross_Morin");
        }, function (error) {

        })
    }

    getCookie();
}]);

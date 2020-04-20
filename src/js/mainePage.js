var ngModule = angular.module('app');

ngModule.controller('mainPageCtrl', ['$scope', '$rootScope', '$http', '$q', '$filter', function ($scope, $rootScope, $http, $q, $filter) {

    GitHubCalendar(".calendar", "MaxwellM", {
        responsive: true,
        proxy: function (url) {
            return "https://maxintosh.org/getGithubInfo/url?url="+ url;
        }
        });

}]);

var ngModule = angular.module('app');

ngModule.controller('mainPageCtrl', ['$scope', '$rootScope', '$http', '$q', '$filter', function ($scope, $rootScope, $http, $q, $filter) {

    GitHubCalendar(".calendar", "maxwellm", {
        responsive: true,
        tooltips: true,
        proxy (username) {
            return $http.get("https://maxintosh.org/getGithubInfo/url?url=https://github.com/users/"+username+"/contributions").then(r => r.data);
        }}
        );
}]);

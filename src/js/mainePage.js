var ngModule = angular.module('app');

ngModule.controller('mainPageCtrl', ['$scope', '$rootScope', '$http', '$q', '$filter', function ($scope, $rootScope, $http, $q, $filter) {
    $scope.totalContributions = 0;

    GitHubCalendar(".calendar", "maxwellm", {
        responsive: true,
        tooltips: true,
        global_stats: true,
            proxy: function (username) {
                return $http.get("https://maxintosh.org/getGithubInfo/url?url=https://github.com/users/" + username + "/contributions").then(r => r.data);
            }}
        );

    // This waits for GitHubCalendar to be done loading...
    angular.element(function () {
        $scope.totalContributions = document.getElementById("contribNum").innerText.split(" ", 1)[0];
        $scope.$apply();
    });
}]);

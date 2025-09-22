var ngModule = angular.module('app');

ngModule.controller('mainPageCtrl', ['$scope', '$rootScope', '$http', '$q', '$filter', function ($scope, $rootScope, $http, $q, $filter) {
    $scope.totalContributions = 0;

    //GitHubCalendar(".calendar", "maxwellm", {
    //    responsive: true,
    //    tooltips: true,
    //    global_stats: true,
    //        proxy: function (username) {
    //            return $http.get("https://maxintosh.org/getGithubInfo/url?url=https://github.com/users/" + username + "/contributions").then(r => r.data);
    //        }}
    //    );

    // This waits for GitHubCalendar to be done loading...
    //angular.element(function () {
    //    checkForContributions();
    //});

    //function sleep(ms) {
    //    return new Promise(resolve => setTimeout(resolve, ms));
    //}

    // Try once a second, for 10 seconds to update our Contribution Number from GitHub!
    //async function checkForContributions() {
    //    const start = new Date();
    //    const future = start.setSeconds(start.getSeconds() + 10);

    //    while (start <= future) {
    //        if (document.getElementById("contribNum") != null && Number(document.getElementById("contribNum").innerText.split(" ", 1)[0]) !== 0){
    //            $scope.totalContributions = document.getElementById("contribNum").innerText.split(" ", 1)[0];
    //            $scope.$apply();
    //            break;
    //        } else {
    //            // Couldn't find it, so lets try again after one second
    //            await sleep(1000);
    //        }
    //    }
    //}
}]);

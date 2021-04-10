var ngModule = angular.module('app');

ngModule.controller('goExampleTwitterCtl', ['$scope', '$http', '$q', '$filter', '$sanitize', function ($scope, $http, $q, $filter, $sanitize) {

    $scope.postTweet = postTweet;
    $scope.getMostRecentTweet = getMostRecentTweet;
    $scope.tweet = "";
    $scope.tweetID = "";

    function postTweet() {
        let obj = {
            Tweet: $scope.tweet
        };
        console.log("Tweet: ", $scope.tweet);
        $http.post("/api/postTweet", obj).then(function (res) {
            let results;
            results = res.data;
            $scope.tweetID = results.id_str;

        }, function(error) {
            alert(error.data);
            //$scope.translation = "";
        });
    }

    function getMostRecentTweet() {
        $http.get("/getTweet").then(function (res) {
            let results;
        }, function (error) {
            alert(error.data);
        })
    }
}]);

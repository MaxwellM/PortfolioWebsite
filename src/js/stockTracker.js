var ngModule = angular.module('app');

ngModule.controller('stockTrackerCtrl', ['$scope', '$http', '$q', '$filter', '$sanitize', '$interval', function ($scope, $http, $q, $filter, $sanitize, $interval) {

    $scope.progressValue = 0;
    $scope.lastUpdate;
    $scope.promise = null;

    let bestBuy = "https://www.bestbuy.com/site/searchpage.jsp?_dyncharset=UTF-8&id=pcat17071&iht=y&keys=keys&ks=960&list=n&qp=category_facet%3Dpcmcat1484077694025&sc=Global&st=nintendo%20switch&type=page&usc=All%20Categories";
    let target = "https://www.target.com/s?sortBy=relevance&Nao=0&category=piakr&searchTerm=nintendo+switch";
    let walmart = "https://www.walmart.com/search/?cat_id=2636_4646529_2002476&facet=retailer%3AWalmart.com";
    let gameStop = "https://www.gamestop.com/video-games/switch/consoles";
    
    function getNewInventory() {
        $scope.progressValue = 0;
        let url = bestBuy;
        // We're going to make multiple get requests, to each vendor. We don't care when they're done, so
        // chaining isn't necessary. We actually want them to complete as fast as possible. Each completed
        // request is going to add progress to our progress bar value. How cool is that?
        $http.get('/getNewInventory/url?url='+ url).then(function (res) {
            let results;
            results = res.data;
            $scope.progressValue += 25;
        }, function(error) {
            alert(error.data);
        });

        url = target;
        $http.get('/getNewInventory/url?url='+ url).then(function (res) {
            let results;
            results = res.data;
            $scope.progressValue += 25;
        }, function(error) {
            alert(error.data);
        });

        url = walmart;
        $http.get('/getNewInventory/url?url='+ url).then(function (res) {
            let results;
            results = res.data;
            $scope.progressValue += 25;
        }, function(error) {
            alert(error.data);
        });

        url = gameStop;
        $http.get('/getNewInventory/url?url='+ url).then(function (res) {
            let results;
            results = res.data;
            $scope.progressValue += 25;
        }, function(error) {
            alert(error.data);
        });



        // Lets update lastUpdate with our new time!
        let today = new Date();
        $scope.lastUpdate = today.getHours() + ":" + today.getMinutes() + ":" + today.getSeconds();
    }

    // starts the interval
    $scope.start = function() {
        // stops any running interval to avoid two intervals running at the same time
        $scope.stop();

        // store the interval promise
        $scope.promise = $interval(getNewInventory, 15000);
    };

    // stops the interval
    $scope.stop = function() {
        $interval.cancel($scope.promise);
    };

    // If this gets destroyed (when leaving the tab) we'll stop the pinging!
    $scope.$on('$destroy', function() {
        console.log("cancelling interval");
        $interval.cancel($scope.promise);
    });

    function parseTarget(html) {


    }

    getNewInventory();
    $scope.start();
}]);

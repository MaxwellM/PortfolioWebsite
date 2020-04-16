var ngModule = angular.module('app');

ngModule.controller('stockTrackerCtrl', ['$scope', '$http', '$q', '$filter', '$sanitize', '$interval', function ($scope, $http, $q, $filter, $sanitize, $interval) {

    $scope.progressValue = 0;
    $scope.lastUpdate;
    $scope.promise = null;
    //$scope.promise2 = null;
    $scope.results = [];

    $scope.items = ['Nintendo Switch'];
    $scope.itemSelected = "";

    let bestBuy = "https://www.bestbuy.com/site/nintendo-switch/nintendo-switch-consoles/pcmcat1484077694025.c?id=pcmcat1484077694025";
    let target = "https://www.target.com/c/nintendo-switch-consoles-video-games/-/N-piakr";
    let walmart = "https://www.walmart.com/search/?cat_id=2636_4646529_2002476&facet=retailer%3AWalmart.com";
    let gameStop = "https://www.gamestop.com/video-games/switch/consoles";

    // function getNewInventorySync() {
    //     $scope.googleData =[];
    //
    //     console.log("START");
    //     $q.all([
    //         $http.get('/getNewInventory/url?url='+ 'https://www.bestbuy.com/site/nintendo-switch/nintendo-switch-consoles/pcmcat1484077694025.c?id=pcmcat1484077694025'),
    //         $http.get('/getNewInventory/url?url='+ 'https://www.target.com/c/nintendo-switch-consoles-video-games/-/N-piakr'),
    //         $http.get('/getNewInventory/url?url='+ 'https://www.walmart.com/search/?cat_id=2636_4646529_2002476&facet=retailer%3AWalmart.com'),
    //         $http.get('/getNewInventory/url?url='+ 'https://www.gamestop.com/video-games/switch/consoles')
    //     ]).then(function(results) {
    //         console.log("END");
    //         /* enter your logic here */
    //         console.log(results);
    //     });
    // }
    
    function getNewInventory() {
        $scope.progressValue = 0;
        $scope.results = [];
        let url = bestBuy;
        let vendor = "BestBuy";
        // We're going to make multiple get requests, to each vendor. We don't care when they're done, so
        // chaining isn't necessary. We actually want them to complete as fast as possible. Each completed
        // request is going to add progress to our progress bar value. How cool is that?
        $http.get('/getNewInventory', {params: {url: url, vendor: vendor}}).then(function (res) {
            let results;
            results = res.data;
            $scope.results.push(results);
            $scope.progressValue += 25;
        }, function(error) {
            alert(error.data);
        });

        url = target;
        vendor = "Target";
        $http.get('/getNewInventory', {params: {url: url, vendor: vendor}}).then(function (res) {
            let results;
            results = res.data;
            $scope.results.push(results);
            $scope.progressValue += 25;
        }, function(error) {
            alert(error.data);
        });

        url = walmart;
        vendor = "Walmart";
        $http.get('/getNewInventory', {params: {url: url, vendor: vendor}}).then(function (res) {
            let results;
            results = res.data;
            $scope.results.push(results);
            $scope.progressValue += 25;
        }, function(error) {
            alert(error.data);
        });

        url = gameStop;
        vendor = "GameStop";
        $http.get('/getNewInventory', {params: {url: url, vendor: vendor}}).then(function (res) {
            let results;
            results = res.data;
            $scope.results.push(results);
            $scope.progressValue += 25;
        }, function(error) {
            alert(error.data);
        });



        // Lets update lastUpdate with our new time!
        let today = new Date();
        // This will get AM/PM to show!
        $scope.lastUpdate= today.toLocaleString('en-US', { hour: 'numeric', minute: 'numeric', second: 'numeric', hour12: true })
    }

    // starts the interval
    $scope.start = function() {
        // stops any running interval to avoid two intervals running at the same time
        $scope.stop();

        // store the interval promise
        // Also run it once every 5 minutes. Don't want to go too crazy.
        $scope.promise = $interval(getNewInventory, 300000);
        //$scope.promise2 = $interval(getNewInventorySync, 15000);
    };

    // stops the interval
    $scope.stop = function() {
        $interval.cancel($scope.promise);
        //$interval.cancel($scope.promise2);
    };

    // If this gets destroyed (when leaving the tab) we'll stop the pinging!
    $scope.$on('$destroy', function() {
        console.log("cancelling interval");
        $interval.cancel($scope.promise);
        //$interval.cancel($scope.promise2);
    });

    function parseTarget(html) {


    }

    getNewInventory();
    $scope.start();
}]);

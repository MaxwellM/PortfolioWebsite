var ngModule = angular.module('app');

ngModule.controller('stockTrackerCtrl', ['$scope', '$http', '$q', '$filter', '$sanitize', '$interval', function ($scope, $http, $q, $filter, $sanitize, $interval) {
    //$scope.getNewInventory = getNewInventory;
    $scope.progressValue = 0;
    $scope.lastUpdate = null;
    $scope.promise = null;
    $scope.getDataPromise = null;
    $scope.results = [];
    $scope.itemSelected = "Nintendo Switch";
    $scope.items = ['Nintendo Switch', 'Dyson V11 Vacuum', 'Apple AirPods', 'XBox Series X', 'XBox Series S'];
    $scope.vendors = [
        {   url: "",
            vendor: "BestBuy"
        },
        {   url: "",
            vendor: "Target"
        },
        {   url: "",
            vendor: "Walmart"
        },
        {   url: "",
            vendor: "GameStop"
        }
    ];

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

    $scope.getDataPromise = function() {
        $scope.myLimit = 10;
        $scope.myPage = 1;
        $scope.progressValue = 0;
        $scope.results = [];
        // Lets loop through all of our vendors and gets some results!
        for(let vendor of $scope.vendors){
            let promise = $http.get('/getNewInventory', {params: {vendor: vendor.vendor, item: $scope.itemSelected}}).then(function (res) {
                let results;
                results = res.data;
                if (results) {
                    $scope.results.push(...results);
                }
                $scope.progressValue += (100 / $scope.vendors.length);
            }, function(error) {
                alert(error.data);
            });
        }
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
        $scope.promise = $interval($scope.getDataPromise, 300000);
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

    function quickWindowResize() {

    }

    $scope.getDataPromise();
    $scope.start();
}]);

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

    $scope.timer = function() {
        let duration = Math.floor((Date.now() - $scope.start)/1000);
        $scope.progressValue = 100 - ((duration / 300)* 100);
        console.log("Percentage: ", $scope.progressValue);
    }

    $scope.getDataPromise = function() {
        //let started = Date.now();
        $interval.cancel($scope.timerInterval);
        $scope.myLimit = 10;
        $scope.myPage = 1;
        $scope.progressValue = 0;
        $scope.results = [];
        // Lets loop through all of our vendors and gets some results!
        const lastElement = $scope.vendors[$scope.vendors.length - 1];
        let returnedPromises = 0;
        for(let vendor of $scope.vendors){
            let promise = $http.get('/getNewInventory', {params: {vendor: vendor.vendor, item: $scope.itemSelected}}).then(function (res) {
                returnedPromises ++;
                let results;
                results = res.data;
                if (results) {
                    $scope.results.push(...results);
                }
                $scope.progressValue += (100 / $scope.vendors.length);
                if (returnedPromises === $scope.vendors.length) {
                    //console.log("SECONDS: ", Math.floor((Date.now() - started) / 1000));
                    //console.log("SECONDS: ", 300 - (Math.floor((Date.now() - started) / 1000)));
                    //timer( 300 - (Date.now() - started));
                    $scope.start = Date.now();
                    $scope.timerInterval = $interval($scope.timer, 1000);
                }
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
        $interval.cancel($scope.timerInterval);
    };

    // If this gets destroyed (when leaving the tab) we'll stop the pinging!
    $scope.$on('$destroy', function() {
        console.log("cancelling interval");
        $interval.cancel($scope.promise);
        $interval.cancel($scope.timerInterval);
    });

    function quickWindowResize() {

    }

    $scope.getDataPromise();
    $scope.start();
}]);

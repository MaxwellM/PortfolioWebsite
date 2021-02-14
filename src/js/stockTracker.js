var ngModule = angular.module('app');

ngModule.controller('stockTrackerCtrl', ['$scope', '$http', '$q', '$filter', '$sanitize', '$interval', function ($scope, $http, $q, $filter, $sanitize, $interval) {
    //$scope.getNewInventory = getNewInventory;
    let fiveMinTime = null;
    $scope.progressValue = 0;
    $scope.lastUpdate = null;
    $scope.promise = null;
    $scope.getDataPromise = null;
    $scope.results = [];
    $scope.itemSelected = "NVIDIA 3070 FE";
    $scope.items = ['Nintendo Switch', 'Dyson V11 Vacuum', 'Apple AirPods', 'XBox Series X', 'XBox Series S', 'NVIDIA 3070 FE', 'NVIDIA 3080 FE'];
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

    // Got this here: https://stackoverflow.com/questions/21294302/converting-milliseconds-to-minutes-and-seconds-with-javascript
    function millisToMinutesAndSeconds(millis) {
        const minutes = Math.floor(millis / 60000);
        const seconds = ((millis % 60000) / 1000).toFixed(0);
        return minutes + ":" + (seconds < 10 ? '0' : '') + seconds;
    }

    $scope.timer = function() {
        let now = Date.now();
        $scope.duration = Math.floor((now - $scope.start)/1000);
        $scope.timeLeft = millisToMinutesAndSeconds(fiveMinTime - now);
        $scope.progressValue = 100 - (($scope.duration / 300)* 100);
        //console.log("Percentage: ", $scope.progressValue);
    }

    function sleep(ms) {
        return new Promise(resolve => setTimeout(resolve, ms));
    }

    $scope.getDataPromise = async function() {
        $interval.cancel($scope.timerInterval);
        $scope.myLimit = 10;
        $scope.myPage = 1;
        $scope.progressValue = 0;
        $scope.progressValueBuffer = 0;
        $scope.currentVendor = "";
        $scope.timeLeft = null;
        $scope.results = [];
        // Lets loop through all of our vendors and gets some results!
        const lastElement = $scope.vendors[$scope.vendors.length - 1];
        let returnedPromises = 0;
        for(let vendor of $scope.vendors){
            $scope.progressValueBuffer += (100 / $scope.vendors.length);
            $scope.currentVendor = vendor.vendor;
            // We can add a sleep, if we want.
            // await sleep(1000);
            let promise = $http.get('/getNewInventory', {params: {vendor: vendor.vendor, item: $scope.itemSelected}}).then(function (res) {
                returnedPromises ++;
                let results;
                results = res.data;
                if (results) {
                    $scope.results.push(...results);
                }
                $scope.progressValue += (100 / $scope.vendors.length);
                // On the last pass, do this...
                if (returnedPromises === $scope.vendors.length) {
                    $scope.start = Date.now();
                    // Five min into the future!
                    fiveMinTime = ($scope.start + 300000);
                    $scope.timerInterval = $interval($scope.timer, 1000);
                    // Lets update lastUpdate with our new time!
                    let today = new Date();
                    $scope.currentVendor = "";
                    $scope.timeLeft = millisToMinutesAndSeconds(fiveMinTime - $scope.start);
                    // This will get AM/PM to show!
                    $scope.lastUpdate = today.toLocaleString('en-US', { hour: 'numeric', minute: 'numeric', second: 'numeric', hour12: true });
                }
            }, function(error) {
                alert(error.data);
            });
        }
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

    $scope.getDataPromise();
    $scope.start();
}]);

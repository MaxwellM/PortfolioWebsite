var ngModule = angular.module('app');

ngModule.controller('unityEggciteBikeCtrl', ['$scope', '$http', '$q', '$filter', function ($scope, $http, $q, $filter) {

    $scope.launchEggciteBike = launchEggciteBike;
    $scope.goBack = goBack;
    $scope.playEggciteBike = false;

    function launchEggciteBike() {
        let gameWindow = document.getElementById("canvas");
        let script = document.createElement("script");
        script.src = "../../unityGames/EggciteBikeWebGL/Release/UnityEggciteBikeLoader.js";
        if (!$scope.playEggciteBike) {
            document.body.appendChild(script);
            $scope.playEggciteBike = true;
        } else {
            //console.log("HERE: ", document.querySelector(script).closest(gameWindow));
            gameWindow.parentNode.removeChild(gameWindow);
            script.remove();
            $scope.playEggciteBike = false;
            // destroy game
            // Module.Quit(function() {
            //     console.log("done!");
            // });
            // Module = null;
        }
    }

    function goBack() {
        window.location.href = '';
    }

    Module = {
        TOTAL_MEMORY: 536870912,
        errorhandler: null,			// arguments: err, url, line. This function must return 'true' if the error is handled, otherwise 'false'
        compatibilitycheck: null,
        backgroundColor: "#222C36",
        splashStyle: "Light",
        dataUrl: "../../unityGames/EggciteBikeWebGL/Release/EggciteBikeWebGL.data",
        codeUrl: "../../unityGames/EggciteBikeWebGL/Release/EggciteBikeWebGL.js",
        asmUrl: "../../unityGames/EggciteBikeWebGL/Release/EggciteBikeWebGL.asm.js",
        memUrl: "../../unityGames/EggciteBikeWebGL/Release/EggciteBikeWebGL.mem",
    };

}]);

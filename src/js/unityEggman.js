var ngModule = angular.module('app');

ngModule.controller('unityEggmanCtrl', ['$scope', '$http', '$q', '$filter', function ($scope, $http, $q, $filter) {

    $scope.launchEggman = launchEggman;
    $scope.goBack = goBack;
    $scope.playEggman = false;

    function launchEggman() {
        let gameWindow = document.getElementById("canvas");
        let script = document.createElement("script");
        script.src = "../../unityGames/EggManWebGL/Release/UnityLoader.js";
        if (!$scope.playEggciteBike) {
            document.body.appendChild(script);
            $scope.playEggman = true;
        } else {
            //console.log("HERE: ", document.querySelector(script).closest(gameWindow));
            gameWindow.parentNode.removeChild(gameWindow);
            script.remove();
            $scope.playEggman = false;
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
        dataUrl: "../../unityGames/EggManWebGL/Release/EggManWebGL.data",
        codeUrl: "../../unityGames/EggManWebGL/Release/EggManWebGL.js",
        //asmUrl: "../../unityGames/EggManWebGL/Release/EggManWebGL.asm.js",
        memUrl: "../../unityGames/EggManWebGL/Release/EggManWebGL.mem",
    };

}]);

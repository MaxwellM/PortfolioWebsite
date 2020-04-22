var ngModule = angular.module('app');

ngModule.controller('unityLostInSpaceCtrl', ['$scope', '$rootScope', '$http', '$q', '$filter', function ($scope, $rootScope, $http, $q, $filter) {

    $scope.launchLostInSpace = launchLostInSpace;
    $scope.goBack = goBack;
    $scope.playLostInSpace = false;
    let script = document.createElement("script");
    function launchLostInSpace() {
        let gameWindow = document.getElementById("canvas");
        script.src = "../../unityGames/LostInSpaceWebGL/Release/UnityLoader.js";
        if (!$scope.playLostInSpace) {
            document.body.appendChild(script);
            $scope.playLostInSpace = true;
        } else {
            //console.log("HERE: ", document.querySelector(script).closest(gameWindow));
            gameWindow.parentNode.removeChild(gameWindow);
            script.remove();
            $scope.playLostInSpace = false;
            // destroy game
            myModule.Quit(function() {
                console.log("done!");
            });
            myModule = null;
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
        dataUrl: "../../unityGames/LostInSpaceWebGL/Release/LostInSpaceWebGL.data",
        codeUrl: "../../unityGames/LostInSpaceWebGL/Release/LostInSpaceWebGL.js",
        asmUrl: "../../unityGames/LostInSpaceWebGL/Release/LostInSpaceWebGL.asm.js",
        memUrl: "../../unityGames/LostInSpaceWebGL/Release/LostInSpaceWebGL.mem",
    };

}]);

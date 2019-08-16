var ngModule = angular.module('app', ['720kb.datepicker']);

ngModule.controller('myCtrl', ['$scope', '$http', '$q', '$filter', function ($scope, $http, $q, $filter) {
    $scope.launchLostInSpace = launchLostInSpace;

    $scope.playLostInSpace = false;

    // Took a while, but found how to load a Unity game on button click, not on page load...
    // https://forum.unity.com/threads/start-unity-player-on-button-click.425180/
    function launchLostInSpace() {
        let gameWindow = document.getElementById("canvas");
        let script = document.createElement("script");
        script.src = "../../unityGames/LostInSpaceWebGL/Release/UnityLoader.js";
        if (!$scope.playLostInSpace) {
            document.body.appendChild(script);
            $scope.playLostInSpace = true;
        } else {
            // destroy game
            //console.log("HERE: ", document.querySelector(script).closest(gameWindow));
            //gameWindow.parentNode.removeChild(gameWindow);
            //$scope.playLostInSpace = false;
        }
    }

}]);

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

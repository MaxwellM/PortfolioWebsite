var ngModule = angular.module('app');

ngModule.controller('angularJSExampleTableCtrl', ['$scope', '$http', '$q', '$filter', function ($scope, $http, $q, $filter) {

    $scope.pushCharacter = pushCharacter;
    $scope.refreshAngularJSExampleTableResults = refreshAngularJSExampleTableResults;
    $scope.setClickedRow = setClickedRow;
    $scope.selectCharacter = selectCharacter;
    $scope.clearSelectedCharacter = clearSelectedCharacter

    $scope.showAddCharacter = false;
    $scope.showViewCharacters = true;

    $scope.searchName = "";
    $scope.searchSpecies = "";
    $scope.allCharacters = [];
    $scope.selectedCharacter = [];

    $scope.query = {
        order: '-name',
        limit: 10,
        page: 1
    };

    $scope.character = {
        id: "",
        name: "",
        homeworld: "",
        born: "",
        died: "",
        species: "",
        gender: "",
        affiliation: "",
        associated: "",
        masters: "",
        apprentices: "",
    };

    function selectCharacter() {
        console.log("Select ID: ", $scope.selectedCharacter[0].id);
        let index = $scope.allCharacters.findIndex(x => x.id === $scope.selectedCharacter[0].id);
        $scope.character = $scope.allCharacters[index];
        console.log("Selected Character: ", $scope.character);
    }

    function clearSelectedCharacter() {
        $scope.character = {
            id: "",
            name: "",
            homeworld: "",
            born: "",
            died: "",
            species: "",
            gender: "",
            affiliation: "",
            associated: "",
            masters: "",
            apprentices: "",
        };
    }

    function setClickedRow(id) {
        let obj = {
            id: id
        };
        console.log("ID: ", obj.id);
        $http.get("/api/setClickedRow", {params: {id}}).then(function (res) {
            let results = res.data;
            console.log("Character: ", results);
            $scope.selectedCharacter = results;
        }, function (error) {
            alert("Couldn't select that Star Wars Character", error);
        })
    }

    function pushCharacter() {
        // update if we're doing that. Otherwise lets add new character
        if ($scope.selectedCharacter.length > 0) {
            let obj = {
                character: $scope.selectedCharacter[0],
            };

            console.log("EDITED CHARACTER: ", $scope.selectedCharacter[0]);

            // Now that everything has been updated, we can update the quote in the DB with the new information!
            $http.post("/api/updateCharacter", obj).then(function (res) {
                let results = res.data;
                refreshAngularJSExampleTableResults();
            }, function (error) {
                alert("Couldn't update the quote!", error);
            })
        } else {
            console.log("Character: ", $scope.character);

            $http.post("/api/addCharacterToDB", $scope.character).then(function (res) {
                let results = res.data;
                Object.keys($scope.character).forEach(k => delete $scope.character[k]);
                refreshAngularJSExampleTableResults();
            }, function(error) {
                alert("ERROR ADDING CHARACTER TO DB: ", error);
            });
        }
    }

    function refreshAngularJSExampleTableResults() {
        let Name = $scope.searchName;
        let Species = $scope.searchSpecies;

        $http.get("/api/loadAngularJSExampleTableResults", {params: {Name, Species}}).then(function (res) {
            let results = res.data;
            $scope.allCharacters = results;
        }, function () {
            console.warn("error loading character results:", arguments);
        });
    }

    function splitResults(results) {
        let split;

        split = results.split(',');

        return split
    }

    // Timed or single shot functions
    refreshAngularJSExampleTableResults();
}]);

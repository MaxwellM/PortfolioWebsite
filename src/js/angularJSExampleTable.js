var ngModule = angular.module('app');

ngModule.controller('angularJSExampleTableCtrl', ['$scope', '$http', '$q', '$filter', function ($scope, $http, $q, $filter) {

    $scope.addCharacter = addCharacter;
    $scope.reSubmitCharacter = reSubmitCharacter;
    $scope.showSection= showSection;
    $scope.refreshAngularJSExampleTableResults = refreshAngularJSExampleTableResults;
    $scope.setClickedRow = setClickedRow;

    $scope.showAddCharacter = false;
    $scope.showViewCharacters = true;

    $scope.searchName = "";
    $scope.searchSpecies = "";
    $scope.allCharacters = [];
    $scope.selectedCharacter = [];

    function setClickedRow(id) {
        let obj = {
            id: id
        };
        console.log("ID: ", obj.id);
        $http.get("/setClickedRow", {params: {id}}).then(function (res) {
            let results = res.data;
            console.log("Character: ", results);
            $scope.selectedCharacter = results;
        }, function (error) {
            alert("Couldn't select that Star Wars Character", error);
        })
    }

    function reSubmitCharacter() {
        let obj = {
            character: $scope.selectedCharacter[0],
        };

        console.log("EDITED CHARACTER: ", $scope.selectedCharacter[0]);

        // Now that everything has been updated, we can update the quote in the DB with the new information!
        $http.post("/updateCharacter", obj).then(function (res) {
            let results = res.data;

        }, function (error) {
            alert("Couldn't update the quote!", error);
        })
    }

    function addCharacter(name, homeworld, born, died, species, gender, affiliation, associated, masters, apprentices) {
        let obj = {
            name: name,
            homeworld: homeworld,
            born: born,
            died: died,
            species: species,
            gender: gender,
            affiliation: affiliation,
            associated: associated,
            masters: masters,
            apprentices: apprentices
        };


        console.log("Character: ", obj);

        $http.post("/addCharacterToDB", obj).then(function (res) {
            let results = res.data;
            refreshAngularJSExampleTableResults();
        }, function(error) {
            alert("ERROR ADDING CHARACTER TO DB: ", error);
        });

    }

    function refreshAngularJSExampleTableResults() {
        let Name = $scope.searchName;
        let Species = $scope.searchSpecies;

        $http.get("/loadAngularJSExampleTableResults", {params: {Name, Species}}).then(function (res) {
            let results = res.data;
            $scope.allCharacters = results;
            // Pre-selects a character.
            $scope.selectedCharacter = results;

            // $scope.allCharacters.forEach(function (res) {
            //     res.born = Date.parse(res.born);
            // });
        }, function () {
            console.warn("error loading character results:", arguments);
            //loadAllSpotterResults();
        });
    }

    function showSection(section) {
        let view = document.getElementById('starWarsCharacterView');
        let add = document.getElementById('starWarsCharacterAdd');
        if (section) {
            $scope.showAddCharacter = true;
            $scope.showViewCharacters = false;
            add.style.backgroundColor = '#52658F';
            add.style.color = '#f5f8fa';
            view.style.backgroundColor = '#f5f8fa';
            view.style.color = '#282e44';
        } else if (!section) {
            $scope.showViewCharacters = true;
            $scope.showAddCharacter = false;
            view.style.backgroundColor = '#52658F';
            view.style.color = '#f5f8fa';
            add.style.backgroundColor = '#f5f8fa';
            add.style.color = '#282e44';
        } else {
            console.log("Not sure what you selected there?");
        }
    }

    function splitResults(results) {
        let split;

        split = results.split(',');

        return split
    }

    // Timed or single shot functions
    refreshAngularJSExampleTableResults();
    showSection(false);
}]);

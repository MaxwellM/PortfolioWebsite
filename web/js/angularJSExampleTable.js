var ngModule = angular.module('app');

ngModule.controller('angularJSExampleTableCtrl', ['$scope', '$http', '$q', '$filter', function ($scope, $http, $q, $filter) {

    $scope.addCharacter = addCharacter;
    $scope.showSection= showSection;
    $scope.refreshAngularJSExampleTableResults = refreshAngularJSExampleTableResults;
    $scope.setClickedRow = setClickedRow;

    $scope.showAddCharacter = false;
    $scope.showViewCharacters = true;

    $scope.searchName = "";
    $scope.searchSpecies = "";
    $scope.searchBorn = "";
    $scope.searchDied = "";
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


    function addCharacter(name, homeworld, born, died, species, gender, affiliation, associated, masters, apprentices) {
        let obj = {
            name: name,
            homeworld: homeworld,
            born: born,
            died: died,
            species: species,
            gender: gender,
            affiliation: splitResults(affiliation),
            associated: splitResults(associated),
            masters: splitResults(masters),
            apprentices: splitResults(apprentices)
        };


        console.log("Character: ", obj);

        $http.post("/addCharacterToDB", obj).then(function (res) {
            let results = res.data;
            console.log("RESULTS: ", results);
        }, function(error) {
            alert("ERROR ADDING CHARACTER TO DB: ", error);
        });

    }

    function refreshAngularJSExampleTableResults() {
        let name = $scope.searchName;
        let species = $scope.searchSpecies;
        let born = Date.parse($scope.searchBorn);
        let died = Date.parse($scope.searchDied);

        $http.get("/loadAngularJSExampleTableResults", {params: {name, species, born, died}}).then(function (res) {
            let results = res.data;

            console.log("CHARACTER RESULTS: ", results);

            $scope.allCharacters = results;

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
        console.log("SECTION: ", section);
        if (section) {
            $scope.showAddCharacter = true;
            $scope.showViewCharacters = false;
            add.style.backgroundColor = '#52658F';
            add.style.color = '#F7F5E6';
            view.style.backgroundColor = '#F7F5E6';
            view.style.color = '#282e44';
        } else if (!section) {
            $scope.showViewCharacters = true;
            $scope.showAddCharacter = false;
            view.style.backgroundColor = '#52658F';
            view.style.color = '#F7F5E6';
            add.style.backgroundColor = '#F7F5E6';
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

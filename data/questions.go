package data

import (
	"strings"

	"github.com/exam-platform/chem-service/models"
	// "github.com/sirupsen/logrus"
)

// var log = logrus.New()

var QuestionBank = []models.Question{
	// Grade 9 - Matter and Its Properties
	{1, "What is the atomic number of Carbon?", []string{"4", "6", "8", "12"}, "B", "Atomic Structure", 9, "easy", "Atoms and Molecules", 1, ""},
	{2, "Water has the chemical formula:", []string{"H₂O₂", "H₂O", "HO", "H₃O"}, "B", "Chemical Compounds", 9, "easy", "Chemical Formulas", 1, ""},
	{3, "Which of the following is a physical change?", []string{"Burning wood", "Rusting iron", "Melting ice", "Souring milk"}, "C", "Matter", 9, "easy", "Physical and Chemical Changes", 1, "Physical changes are reversible"},
	{4, "The number of protons in an atom is equal to its:", []string{"Mass number", "Atomic number", "Neutron number", "Electron number"}, "B", "Atomic Structure", 9, "easy", "Atomic Structure", 1, ""},
	{5, "Sodium chloride (NaCl) is an example of:", []string{"Covalent compound", "Ionic compound", "Metallic compound", "Mixture"}, "B", "Chemical Bonding", 9, "medium", "Chemical Bonding", 2, ""},

	// Grade 9 - Periodic Table
	{6, "The periodic table is arranged in order of increasing:", []string{"Atomic mass", "Atomic radius", "Atomic number", "Electronegativity"}, "C", "Periodic Table", 9, "easy", "Periodic Classification", 1, ""},
	{7, "Which element is represented by the symbol 'Fe'?", []string{"Fluorine", "Francium", "Iron", "Fermium"}, "C", "Periodic Table", 9, "easy", "Elements", 1, "Fe comes from Latin 'Ferrum'"},
	{8, "Noble gases belong to which group?", []string{"Group 1", "Group 17", "Group 18", "Group 16"}, "C", "Periodic Table", 9, "easy", "Groups and Periods", 1, ""},

	// Grade 10 - Chemical Reactions
	{9, "In the reaction 2H₂ + O₂ → 2H₂O, the ratio of H₂ to O₂ is:", []string{"1:1", "2:1", "1:2", "3:1"}, "B", "Chemical Reactions", 10, "medium", "Stoichiometry", 2, ""},
	{10, "What type of reaction is: A + BC → AC + B?", []string{"Combination", "Decomposition", "Displacement", "Double displacement"}, "C", "Chemical Reactions", 10, "easy", "Types of Reactions", 1, ""},
	{11, "The pH of a neutral solution at 25°C is:", []string{"0", "7", "14", "1"}, "B", "Acids and Bases", 10, "easy", "pH Scale", 1, ""},
	{12, "What is produced when an acid reacts with a base?", []string{"Acid + Base", "Salt + Water", "Gas + Precipitate", "Oxide + Hydrogen"}, "B", "Acids and Bases", 10, "easy", "Neutralization", 1, ""},

	// Grade 10 - Mole Concept
	{13, "Avogadro's number is approximately:", []string{"6.022 × 10²³", "6.022 × 10²²", "3.14 × 10²³", "1.6 × 10¹⁹"}, "A", "Mole Concept", 10, "easy", "Mole Concept", 1, ""},
	{14, "The molar mass of NaCl is (Na=23, Cl=35.5):", []string{"58", "58.5", "68.5", "48.5"}, "B", "Mole Concept", 10, "medium", "Molar Mass", 2, "23 + 35.5 = 58.5 g/mol"},
	{15, "How many moles are in 18 g of water? (Molar mass H₂O = 18 g/mol)", []string{"2 mol", "18 mol", "1 mol", "0.5 mol"}, "C", "Mole Concept", 10, "easy", "Mole Calculations", 1, ""},

	// Grade 11 - Thermodynamics
	{16, "The enthalpy change for a combustion reaction is always:", []string{"Positive (endothermic)", "Negative (exothermic)", "Zero", "Variable"}, "B", "Thermodynamics", 11, "easy", "Enthalpy", 1, ""},
	{17, "Hess's Law states that:", []string{"Enthalpy change depends on path", "Enthalpy change is independent of path", "All reactions are exothermic", "Temperature affects enthalpy"}, "B", "Thermodynamics", 11, "medium", "Thermochemistry", 2, ""},
	{18, "The standard entropy change ΔS for a spontaneous process must be:", []string{"Always positive", "Always negative", "Zero", "The total ΔS (system + surroundings) must be positive"}, "D", "Thermodynamics", 11, "hard", "Entropy", 3, "Second law of thermodynamics"},

	// Grade 11 - Chemical Equilibrium
	{19, "Le Chatelier's Principle states that when a system at equilibrium is disturbed:", []string{"It moves to a new equilibrium", "It restores itself to oppose the change", "The equilibrium constant changes", "The reaction stops"}, "B", "Equilibrium", 11, "medium", "Chemical Equilibrium", 2, ""},
	{20, "For the reaction N₂ + 3H₂ ⇌ 2NH₃, increasing pressure will:", []string{"Shift equilibrium left", "Shift equilibrium right", "Have no effect", "Stop the reaction"}, "B", "Equilibrium", 11, "medium", "Le Chatelier's Principle", 2, "Fewer moles of gas on right side"},
	{21, "The equilibrium constant Kc for a reaction depends on:", []string{"Pressure", "Catalyst presence", "Temperature only", "Concentration of reactants"}, "C", "Equilibrium", 11, "hard", "Equilibrium Constant", 3, ""},

	// Grade 11 - Electrochemistry
	{22, "In electrolysis of water, which gas is produced at the cathode?", []string{"Oxygen", "Hydrogen", "Chlorine", "Nitrogen"}, "B", "Electrochemistry", 11, "easy", "Electrolysis", 1, "Cathode: reduction occurs"},
	{23, "The standard electrode potential of the hydrogen electrode is:", []string{"+1.23 V", "-0.76 V", "0.00 V", "+0.34 V"}, "C", "Electrochemistry", 11, "medium", "Electrode Potentials", 2, "Reference electrode by definition"},

	// Grade 12 - Organic Chemistry
	{24, "The IUPAC name for CH₃-CH₂-CH₂-OH is:", []string{"Methanol", "Ethanol", "Propan-1-ol", "Butan-1-ol"}, "C", "Organic Chemistry", 12, "medium", "Nomenclature", 2, "3 carbons with -OH at position 1"},
	{25, "What type of reaction occurs when CH₄ reacts with Cl₂ in presence of UV light?", []string{"Addition", "Substitution", "Elimination", "Polymerization"}, "B", "Organic Chemistry", 12, "medium", "Organic Reactions", 2, "Free radical substitution"},
	{26, "Benzene is an aromatic compound because:", []string{"It has alternating single and double bonds", "It follows Hückel's rule (4n+2 π electrons)", "It is cyclic", "It contains only carbon and hydrogen"}, "B", "Organic Chemistry", 12, "hard", "Aromaticity", 3, ""},

	// Grade 12 - Polymers and Biomolecules
	{27, "Nylon-6,6 is formed from:", []string{"Adipic acid + Hexamethylenediamine", "Caprolactam", "Ethylene", "Vinyl chloride"}, "A", "Polymers", 12, "hard", "Polymers", 3, "Condensation polymerization"},
	{28, "DNA is made up of:", []string{"Amino acids", "Nucleotides", "Fatty acids", "Monosaccharides"}, "B", "Biomolecules", 12, "easy", "Biomolecules", 1, ""},
	{29, "Which vitamin is synthesized in the skin upon exposure to sunlight?", []string{"Vitamin A", "Vitamin B12", "Vitamin C", "Vitamin D"}, "D", "Biomolecules", 12, "easy", "Vitamins", 1, ""},
	{30, "The catalyst used in the Haber process for ammonia synthesis is:", []string{"Platinum", "Vanadium pentoxide", "Iron with K₂O and Al₂O₃", "Nickel"}, "C", "Industrial Chemistry", 12, "hard", "Industrial Processes", 3, ""},
}

// GetQuestionsByGrade filters questions by grade
func GetQuestionsByGrade(grade, n int) []models.Question {
	// log.Debugf("Querying question bank for grade=%d, limit=%d", grade, n)
	var result []models.Question
	for _, q := range QuestionBank {
		if q.Grade == grade {
			result = append(result, q)
			if len(result) >= n {
				break
			}
		}
	}
	// log.Infof("Retrieved %d questions for grade %d", len(result), grade)
	return result
}

// GetQuestionsByTopic filters questions by topic
func GetQuestionsByTopic(topic string, n int) []models.Question {
	// log.Debugf("Querying question bank for topic='%s', limit=%d", topic, n)
	var result []models.Question
	for _, q := range QuestionBank {
		if strings.EqualFold(q.Topic, topic) {
			result = append(result, q)
			if len(result) >= n {
				break
			}
		}
	}
	// log.Infof("Retrieved %d questions for topic '%s'", len(result), topic)
	return result
}

// GetQuestionsByComplexity filters questions by complexity
func GetQuestionsByComplexity(complexity string, n int) []models.Question {
	// log.Debugf("Querying question bank for complexity='%s', limit=%d", complexity, n)
	var result []models.Question
	for _, q := range QuestionBank {
		if strings.EqualFold(q.Complexity, complexity) {
			result = append(result, q)
			if len(result) >= n {
				break
			}
		}
	}
	// log.Infof("Retrieved %d questions with complexity '%s'", len(result), complexity)
	return result
}

// GetAllTopics returns unique topics from the question bank
func GetAllTopics() []string {
	// log.Debug("Fetching all available topics from question bank")
	seen := make(map[string]bool)
	var topics []string
	for _, q := range QuestionBank {
		if !seen[q.Topic] {
			seen[q.Topic] = true
			topics = append(topics, q.Topic)
		}
	}
	// log.Infof("Found %d unique topics", len(topics))
	return topics
}

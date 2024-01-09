package data

import "example.com/internal/models"

var SwmsSchemaData = []models.SwmsSchema{
	{
		ID:               1,
		SubId:            1,
		Name:             "Arrival at site. Unloading and Set-Up.",
		Task:             []string{"Unload vehicle"},
		PotentialHazards: []string{"Musculoskeletal strains", "Slips, trips and falls"},
		RiskBefore:       "3",
		ControlMeasures: []models.ControlMeasure{
			{
				Name: "Planning, Consultation, Adherence to Manual Handling Techniques",
				Values: []string{
					"When unloading the vehicle we will ensure that we are as close as possible to the area where the equipment will be set up. If required we will seek out assistance in unloading heavy items, however our normal work does not include heavy items.",
					"We will use sensible manual handling techniques making sure our backs are straight and bending with the knees.",
				},
			},
		},
		RiskAfter: "5",
	},
	{
		ID:               1,
		SubId:            2,
		Name:             "Arrival at site. Unloading and Set-Up.",
		Task:             []string{"Working in the sun Dangerous UV Rays"},
		PotentialHazards: []string{"Exposure to UV radiation.", "Heat stress", "De-hydration", "Collapse", "Nauseated", "Skin Cancer", "Bodily Injury", "Infection", "Death"},
		RiskBefore:       "1",
		ControlMeasures: []models.ControlMeasure{
			{
				Name: "Planning and Consultation",
				Values: []string{
					"Work health and safety legislation in each Australian state requires your employer or PCBU (person conducting a business undertaking) to provide a safe working environment.",
					"Skin cancer is a preventable disease and will actively promote, encourage and support skin protection in all work activities with which they are associated.",
					"All employees or Contractors must wear clothing to protect from the harmful UV Rays.",
					"Best options to avoid skin cancer when working outside",
					"Shirts or tops which have longer sleeves and a collar.",
					"Longer legged shorts where appropriate.",
					"Wide brimmed or legionnaire hats whenever practical.",
					"Eye protection tinted safety glasses.",
					"Actively encourage all employees to routinely apply broad spectrum water resistant 30+ sunscreen and stress the importance of regular re-application.",
					"Advise all workers, about the UV Protection Policy and encourage them to comply with it.",
					"Work and take breaks in the shade. Where no shade exists, use temporary portable shade.",
					"If possible, Plan to work indoors or in the shade during the middle of the day when UV radiation levels are strongest.",
					"Plan to do outdoor work tasks early in the morning or later in the afternoon when UV radiation levels are lower.",
					"Share outdoor tasks and rotate staff so the same person is not always out in the sun.",
					"Choose shade that has extensive overhead and side cover and is positioned away from highly reflective surfaces.",
				},
			},
		},
		RiskAfter: "2",
	},
	{
		ID:               1,
		SubId:            3,
		Name:             "Arrival at site. Unloading and Set-Up.",
		Task:             []string{"Unloading vehicle (cont.)"},
		PotentialHazards: []string{"Electrical Hazards", "Fire"},
		RiskBefore:       "1",
		ControlMeasures: []models.ControlMeasure{
			{
				Name: "Risk Assessment, Planning and Consultation",
				Values: []string{
					"Before commencing any work in the roof we will consider whether live electrical wiring is a hazard.",
					"If live electrical wiring is a hazard we will consider cutting the house power and using an independent power source such as generator or neighbours power.",
					"We will walk through the premises with the owner to identify the location of all down lights and other ceiling accessories.",
					"We will record the location and type and then make the necessary precautions when laying the insulation. As a default we will leave a clearance of 50mm from incandescent lights and 200mm from halogen lights including 50mm for any transformer, unless the lights are fitted with a suitable fire rated enclosure.",
				},
			},
		},
		RiskAfter: "5",
	},
	{
		ID:               2,
		SubId:            1,
		Name:             "General Construction",
		Task:             []string{"Use of hand and power tools"},
		PotentialHazards: []string{"Electrocution", "Cuts and abrasions", "Eye and hearing damage"},
		RiskBefore:       "1",
		ControlMeasures: []models.ControlMeasure{
			{
				Name: "Safety Glasses, Ear Protection and RCD.",
			},
			{
				Name: "All Electrical leads and tools will be tested and tagged every 3 months in accordance with AS/NZS 3012:2010. A test register will also be available for inspection",
			},
			{
				Name: "Guards on tools and equipment will be maintained and working effectively before being used on site.",
			},
			{
				Name: "Guarding on tools will not be removed to perform any work activity.",
			},
		},
		RiskAfter: "4",
	},
	{
		ID:               2,
		SubId:            2,
		Name:             "General Construction",
		Task:             []string{"Use of hand and power tools (cont.)"},
		PotentialHazards: []string{"Exposure to UV radiation.", "Heat stress", "De-hydration", "Collapse", "Nauseated", "Skin Cancer", "Bodily Injury", "Infection", "Death"},
		RiskBefore:       "2",
		ControlMeasures: []models.ControlMeasure{
			{
				Name: "All tools and equipment will be inspected prior to work activity for any faults or defects. If a fault or defect is found the item will be removed from services, and reported to the supervisor as soon as practicable.",
			},
			{
				Name: "All persons performing work where there is a risk of a foreign object striking the eye, should consider wearing eye protection. If an item of plant or equipment creates excessive noise, that is where you need to raise your voice to talk, we will wear appropriate hearing protection and if there is a risk of injury to the head by falling objects then we will wear hard hats.",
			},
			{
				Name: "When we use plant, equipment or power tools we will also follow the manufacturer’s instructions for the correct PPE to be worn and the safe use instructions. We will be competent in the use of the PPE and risk assessments must be undertaken prior to using PPE to show that the hierarchy of control was used in determining whether or not to use PPE.",
			},
		},
		RiskAfter: "2",
	},
	{
		ID:               2,
		SubId:            3,
		Name:             "General Construction",
		Task:             []string{"Using Ladders"},
		PotentialHazards: []string{"Falling"},
		RiskBefore:       "1",
		ControlMeasures: []models.ControlMeasure{
			{
				Name: "Tie Offs, Base Support, Gutter Anchors, Levellers",
				Values: []string{
					"All ladders used on site will be rated ‘Industrial’ with 120kg (minimum) load rating. A single and extension ladders must be secured at the top, bottom or both. Persons using the ladder must have 3 points of contact at all times (i.e. 2 hands and 1 foot or 2 feet and 1 hand or be holding a stable object e.g. gutter, wall frame). Ladders are to be maintained in a sound working condition and be appropriate for the task to be undertaken. Tools requiring two handed operation or a high degree of leverage force should not be used while on ladders. A ladder is not a work platform.",
				},
			},
		},
		RiskAfter: "4",
	},
	{
		ID:               2,
		SubId:            4,
		Name:             "General Construction",
		Task:             []string{"Sweeping"},
		PotentialHazards: []string{"Dust – silicosis (RCS)", ""},
		RiskBefore:       "1",
		ControlMeasures: []models.ControlMeasure{
			{
				Name: "Dust Mask, Eye Protection, Wet Down Area",
				Values: []string{
					"We will assess whether to wet down areas to reduce dust emission form works conducted. Where the risk of dust production is high, worker will wear appropriate PPE and refer to Engineering Controls that will reduce Silica Dust exposure.",
				},
			},
			{
				Name: "RCS dust should not be disturbed by use of compressed air, blowers or sweeping.",
			},
			{
				Name: "Training Consultation & Supervision",
				Values: []string{
					"Frequent job rotation",
					"Avoid twisting",
					"Correct posture at all times",
					"Use electric floor sweeper where possible",
				},
			},
		},
		RiskAfter: "4",
	},
	{
		ID:               3,
		SubId:            1,
		Name:             "Working with Silica",
		Task:             []string{"Concrete Floor Grinding", "Concrete Cutting", "Removal & cutting wall/Floor Tiles.", "Sanding Plaster Board", "Grinding Villa Board", "Cutting", "Grinding Masonry Bricks/Blocks"},
		PotentialHazards: []string{"Dust – silicosis (RCS)", "Lung cancer", "Chronic obstructive pulmonary disease", "Kidney disease"},
		RiskBefore:       "1",
		ControlMeasures: []models.ControlMeasure{
			{
				Name: "Where possible, work will be undertaken off-site. (such as pre-cutting to size, pre-drilling etc)",
			},
			{
				Name: "Relevant safety data sheet (SDS) will be obtained for products. If silica presence is uncertain, will assume it is.",
				Values: []string{
					"All workers must familiarise themselves with the information supplied on the safety data sheet (SDS) that silica is likely to be present and comply with the requirements within.",
				},
			},
			// ... [Continue with all ControlMeasures as in old schema]
		},
		RiskAfter: "4",
	},
	{
		ID:               3,
		SubId:            2,
		Name:             "Working with Silica",
		Task:             []string{"Concrete Floor Grinding", "Concrete Cutting", "Removal & cutting wall/Floor Tiles.", "Sanding Plaster Board", "Grinding Villa Board", "Cutting", "Grinding Masonry Bricks/Blocks (Cont.)"},
		PotentialHazards: []string{"Exposure to UV radiation.", "Heat stress", "De-hydration", "Collapse", "Nauseated", "Skin Cancer", "Bodily Injury", "Infection", "Death"},
		RiskBefore:       "2",
		ControlMeasures:  []models.ControlMeasure{
			// ... [Include all ControlMeasures as in old schema]
		},
		RiskAfter: "2",
	},
	{
		ID:               4,
		SubId:            1,
		Name:             "Manual Handling",
		Task:             []string{"Manual handling / locations of the loads and distances to be moved"},
		PotentialHazards: []string{"Back, shoulder strain", "Fatigue"},
		RiskBefore:       "3",
		ControlMeasures: []models.ControlMeasure{
			{
				Name: "Training Consultation & Supervision",
				Values: []string{
					"Use mechanical handling equipment",
					"Team lifting",
					"Modify work place layout so materials will not be carried long distances",
					"Ensure clear access and egress",
				},
			},
		},
		RiskAfter: "5",
	},
	{
		ID:               5,
		SubId:            1,
		Name:             "Asbestos Removal",
		Task:             []string{"Sheeting and guttering"},
		PotentialHazards: []string{"Asbestos related diseases"},
		RiskBefore:       "1",
		ControlMeasures: []models.ControlMeasure{
			{
				Name: "Monitoring, Supervision, Training, PPE, Specialised Equipment.",
				Values: []string{
					"All workers directly involved with the removal, and or handling of Asbestos will hold a general safety induction card and an approved Bonded Asbestos Removal Certificate, issued by Queensland WHS.",
					"Only workers directly involved with the removal will be present in the area where the removal is taking place.",
					"Signage and barriers will be erected if other persons are present. All workers involved in the removal will wear P2 disposable respirators (masks) and disposable coveralls.",
				},
			},
			{
				Name: "All asbestos sheeting and gutters will be removed in full pieces where possible. Nails will be punched and screws removed, along with any trims holding the sheets in position.",
			},
			{
				Name: "Power tools will not be used on the sheeting or gutters and no cutting will take place.",
			},
			{
				Name: "External sheeting and gutters will be wet down prior to removal. Roof sheeting will not be wet down prior to removal as it will create a slip hazard and put the workers at risk of an injury. Any internal sheeting will already be sealed by existing paint, wetting down would be of no benefit and would cause damage to the floors and ceilings.",
			},
			{
				Name: "Once the internal sheeting is removed the area will be vacuumed with an industrial vacuum fitted with a HEPA filter. The vacuum bags will also be placed in the 200 micro metre polythene bags and disposed of. On completion of the decontamination the area will be able to be accessed by persons who were not directly involved with the removal.",
			},
			{
				Name: "Workers will wash any exposed parts of their body i.e. face and hands before stopping for morning tea or lunch to eat and before leaving site.",
			},
		},
		RiskAfter: "4",
	},
	{
		ID:               5,
		SubId:            2,
		Name:             "Asbestos Removal",
		Task:             []string{"Bonded or friable asbestos in excess of 10 sq. metres."},
		PotentialHazards: []string{"Asbestos related diseases"},
		RiskBefore:       "1",
		ControlMeasures: []models.ControlMeasure{
			{
				Name: "Monitoring, Supervision, training, PPE, Specialised Equipment.",
				Values: []string{
					"A competent person will supervise the Asbestos removal work at all times whilst the work is being undertaken.",
				},
			},
			{
				Name: "All workers will hold a general induction card. Only workers directly involved with the removal will be present in the area where the removal is taking place. Signage and barriers will be erected if other persons are present. All workers involved in the removal will wear P2 disposable respirators (masks) and disposable coveralls and gloves.",
			},
			{
				Name: "The ACM will be removed using wet methods and contained within an enclosed area.",
			},
			{
				Name: "All ventilation and Air-conditioning Networks servicing the ACM area will be closed down for the duration of the work and all vents sealed to prevent entry of airborne asbestos fibres into ducts.",
			},
			{
				Name: "After work ceases all ventilation filters for recirculated air will be replaced.",
			},
			{
				Name: "We will take care not to allow asbestos fibres to escape via pipe or conduit holes.",
			},
			{
				Name: "We shall establish a negative pressure work area for the removal of the ACM and this area will be set up in accordance with the provisions of the Code of Practice for the Safe removal of Asbestos 2nd edition. [NOHSC:2002(2005)] Latest Version 2018.",
			},
			{
				Name: "We will only use grinding or abrading tools where no other alternative is available and only after a written risk assessment has been undertaken.",
			},
			{
				Name: "We will set up and use an on-site decontamination unit.",
			},
			{
				Name: "We are aware of and will enforce 'No laundering of contaminated protective clothing in workers’ homes'.",
			},
			{
				Name: "On completion of the work a competent person, other than the works supervisor, will conduct a site clearance and will issue a clearance certificate.",
			},
		},
		RiskAfter: "4",
	},
}

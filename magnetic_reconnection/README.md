Bz Component Of IMF 

When the interplanetary magnetic field has a large negative (southward) Bz component, interconnection between the interplanetary and terrestrial
magnetic fields is greatly enhanced, with a resulting release of energy. Prolonged periods of southward magnetic fields often result in geomagnetic storms, with
associated auroral displays and space weather effects on communications and space hardware.

When there is a temporary disturbance of earth’s magnetosphere, then a geomagnetic storm (or solar storm) has occurred. It is caused by a solar wind shock wave 
and/or cloud of a magnetic field that interacts with the Earth's magnetic fields. The interaction of Interplanetary Magnetic Fields (IMF) of the sun with the 
earth’s magnetic fields in opposite directions is known as magnetic reconnection. Magnetic reconnection is the breaking
and reconnecting of oppositely directed magnetic field lines in plasma at a neutral point which leads to converting the magnetic field energy into plasma kinetic and thermal energy. 

The more intense the geostorms, the more Bz would go south and the more
magnetic reconnection and subsequently auroral substorms which can cause disruption of shortwave radio communications, distortion
of compass readings in polar regions, failure of electrical transmission lines, increased corrosion in
long pipelines, anomalies in the operations of communications satellites, and potentially lethal
doses of radiation for astronauts in interplanetary spacecrafts.


KP Index

The Kp index is used to characterize the magnitude of geomagnetic storms. Kp is an excellent indicator of disturbances in the Earth's magnetic field. The principal users affected
by geomagnetic storms are the electrical power grid, spacecraft operations, users of radio signals that reflect off of or pass through the ionosphere, and observers of the aurora.


Project Objectives

(notebook: magnetic_recon.ipynb)
-Using Bz data recorded by ACE, DSCOVR and WIND spacecrafts to determine the frequency of magnetic reconnection occurences in a given year.
-Rendering Bz vs. Time plots based on user input spacecraft name (ACE/WIND/DSCOVR) and user input year.
-Applying wavelet decomposition on each graph to carry out wave denoising (main purpose is to reduce unwanted noise and reconstruct the signal)

(notebook: aurora_visuals.ipynb)
-Fetching data from the NOAA daily aurora forecast in order to create kml files that can be imported in Google Earth.
-These files contain data regarding the longitude, latitude and Kp index values (to determine auroral intensity) which will be used to pin locations on Google Earth where auroras are most likely to be seen.
-3 files will be generated: ovation_aurora_north.kml (for aurora borealis), ovation_aurora_south.kml(for aurora australis) and ovation_aurora_latest_max.kml (for the aurora with the maximum Kp index/ intensity)
-To view the locations on Google Earth, download the files and import them in Google Earth.
